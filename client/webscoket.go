package client

import (
	"log"
	"time"
	"crypto/tls"
	"net/http"
	"net/url"
	"sync/atomic"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type WSCallback func(conn *websocket.Conn, data interface{}) (error)

type WSClient struct {
	readBufSize        int
	writeBufSize       int
	pingInterval       int
	pingTimeout        int
	retry              int
	retryMax           int
	retryWait          int
	finisable          uint32
	finishRequestChan  chan int
	finishResponseChan chan int
}

type WSRequest struct {
        URL                 string
        Headers             map[string]string
        parsedURL           *url.URL
}

type pingContext struct {
	conn                   *websocket.Conn
	pingFinishRequestChan  chan int
	pingFinishResponseChan chan int
}

func (w *WSClient) messageLoop(conn *websocket.Conn, callback WSCallback, callbackData interface{}) (bool) {
	atomic.StoreUint32(&w.finisable, 1)
	for {
		select {
		case <- w.finishRequestChan:
			return true
		default:
			err := callback(conn, callbackData)
			if err != nil {
				log.Printf("callback error (reason = %v)", err)
				return false
			}
		}
	}
}

func (w *WSClient) pingLoop(pingCtx *pingContext) {
	for {
		select {
		case <-pingCtx.pingFinishRequestChan:
			close(pingCtx.pingFinishResponseChan)
			return
		case <-time.After(time.Duration(w.pingInterval) * time.Second):
			deadline := time.Now()
			deadline.Add(time.Duration(w.pingTimeout) * time.Second)
			pingCtx.conn.WriteControl(websocket.PingMessage, []byte("ping"), deadline)
		}
	}
}

func (w *WSClient) startPing(conn *websocket.Conn) (*pingContext) {
	pingCtx := &pingContext {
		conn: conn,
		pingFinishRequestChan: make(chan int),
		pingFinishResponseChan: make(chan int),
	}
	go w.pingLoop(pingCtx)
	return pingCtx
}

func (w *WSClient) stopPing(pingCtx *pingContext) {
	close(pingCtx.pingFinishRequestChan)
	<-pingCtx.pingFinishResponseChan
	return
}

func  (w *WSClient) connect(request *WSRequest, callback WSCallback, callbackData interface{}, header http.Header, dialer *websocket.Dialer) (bool) {
	conn, response, err := dialer.Dial(request.URL, header)
	if err != nil {
		log.Printf("can not dial (url = %v, reason = %v)", request.URL, err)
		time.Sleep(time.Duration(w.retryWait) * time.Second)
		w.retry += 1
		if w.retry > w.retryMax {
			log.Printf("give up retry (url = %v)", request.URL)
			return false
		}
		return true
	}
	defer conn.Close()
	if response.StatusCode < 200 && response.StatusCode >= 300 {
		log.Printf("error status code (url = %v, status code == %v)", request.URL, response.StatusCode)
		time.Sleep(time.Duration(w.retryWait) * time.Second)
		w.retry += 1
		if w.retry > w.retryMax {
			log.Printf("give up retry (url = %v)", request.URL)
			return false
		}
		return true
	}
	w.retry = 0
	pingContext := w.startPing(conn)
	finish := w.messageLoop(conn, callback, callbackData)
	w.stopPing(pingContext)
	if !finish {
		time.Sleep(time.Duration(w.retryWait) * time.Second)
		return true
	}
	return false
}

func (w *WSClient) connectLoop(request *WSRequest, callback WSCallback, callbackData interface{}) {
	header := http.Header{}
	for k, v := range request.Headers {
		header.Set(k, v)
	}
	dialer := &websocket.Dialer{
		Proxy:           http.ProxyFromEnvironment,
		TLSClientConfig: &tls.Config{ServerName: request.parsedURL.Host},
	}
	for {
		retryable := w.connect(request, callback, callbackData, header, dialer)
		if retryable {
			continue
		}
		close(w.finishResponseChan)
		break
	}
}

func (w *WSClient) Start(request *WSRequest, callback WSCallback, callbackData interface{}) (error) {
	parsedURL, err := url.Parse(request.URL)
	if err != nil {
		return errors.Wrapf(err, "can not parse url (url = %v)", request.URL)
	}
	request.parsedURL = parsedURL
	go w.connectLoop(request, callback, callbackData)
	return nil
}

func (w *WSClient) Stop() {
	if atomic.LoadUint32(&w.finisable) == 0 {
		close(w.finishRequestChan)
		return
	}
	close(w.finishRequestChan)
	<-w.finishResponseChan
}

func NewWSClient(readBufSize int, writeBufSize int, retryMax int, retryWait int) *WSClient {
	if readBufSize == 0 {
		readBufSize = 1024 * 1024 * 2
	}
	if writeBufSize == 0 {
		writeBufSize = 1024 * 1024 * 2
	}
	return &WSClient{
		readBufSize:           readBufSize,
		writeBufSize:          writeBufSize,
		pingInterval:          5,
		pingTimeout:           10,
		retry:                 0,
		retryMax:              retryMax,
		retryWait:             retryWait,
		finisable:             0,
		finishRequestChan:     make(chan int),
		finishResponseChan:    make(chan int),
	}
}
