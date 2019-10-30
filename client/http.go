package client

import (
	"log"
	"fmt"
	"bytes"
	"time"
	"strings"
	"sync"
	"net"
	"net/http"
	"net/url"
	"crypto/tls"
	"io/ioutil"
	"github.com/pkg/errors"
	"github.com/viki-org/dnscache"
)

type RequestMethod int

type HTTPRequest struct {
	PathQuery string
	URL       string
	Method    string
	Headers   map[string]string
	Body      []byte
	parsedURL *url.URL
}

func (h *HTTPRequest)ToString() (string) {
	return fmt.Sprintf("%v %v", h.Method, h.URL)
}

type HTTPClient struct {
	timeout           int
	idleConnTimeout   int
	localAddr         *net.TCPAddr
	localAddrIp       net.IP
	resolver          *dnscache.Resolver
	resolverIdx       int
	resolverIdxMutex  *sync.Mutex
	clientsCache      map[string]*http.Client
	clientsCacheMutex *sync.Mutex
}

func (c *HTTPClient) newHTTPTransport(scheme string, host string) (*http.Transport) {
	newTransport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: func(network string, address string) (net.Conn, error) {
			separator := strings.LastIndex(address, ":")
			ips, _ := c.resolver.Fetch(address[:separator])
			var ip net.IP
			c.resolverIdxMutex.Lock()
			defer c.resolverIdxMutex.Unlock()
			if c.resolverIdx >= len(ips) {
				c.resolverIdx = 0
			}
			var i int
			for i = c.resolverIdx; i < len(ips); i++ {
				tmpIp := ips[i]
				srcIpStr := c.localAddrIp.String()
				dstIpStr := tmpIp.String()
				if (!strings.Contains(srcIpStr, ":") && strings.Contains(dstIpStr, ":")) ||
				   (strings.Contains(srcIpStr, ":") && !strings.Contains(dstIpStr, ":")) {
					continue
				}
				ip = tmpIp
				break
			}
			if ip == nil {
				log.Printf("can not look up address %v in dns cache", address[:separator])
				return net.Dial(network, address)
			}
			c.resolverIdx = i + 1
			return (&net.Dialer{
				LocalAddr: c.localAddr,
				Timeout:   time.Duration(c.timeout) * time.Second,
				KeepAlive: time.Duration(c.timeout) * time.Second,
				//DualStack: true,
			}).Dial("tcp", net.JoinHostPort(ip.String(), address[separator+1:]))
		},
		//ForceAttemptHTTP2: true,
		ResponseHeaderTimeout: time.Duration(c.timeout) * time.Second,
                TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConns:          0,
		MaxIdleConnsPerHost:   100,
		MaxConnsPerHost:       0,
		IdleConnTimeout:       90 * time.Second,
	}
	if scheme == "https" {
		newTransport.TLSClientConfig = &tls.Config{ServerName: host}
	}
	return newTransport
}

func (c *HTTPClient) newClient(scheme string, host string) (*http.Client) {
	c.clientsCacheMutex.Lock()
	defer c.clientsCacheMutex.Unlock()
        clientId := fmt.Sprintf("%v,%v,%v", scheme, host)
	cachedHttpClient, ok := c.clientsCache[clientId]
	if ok {
		return cachedHttpClient
	}
	transport := c.newHTTPTransport(scheme, host)
	newHttpClient := &http.Client{
		Transport: transport,
		Timeout: time.Duration(c.timeout) * time.Second,
	}
	c.clientsCache[clientId] = newHttpClient
	return newHttpClient
}

func (c *HTTPClient) DoRequest(request *HTTPRequest) (*http.Response, []byte, error) {
	parsedURL, err := url.Parse(request.URL)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not parse url (url = %v)", request.URL)
	}
	client := c.newClient(parsedURL.Scheme, parsedURL.Host)
	req, err := http.NewRequest(request.Method, request.URL, bytes.NewBuffer(request.Body))
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create request (method = %v, url = %v, request body = %v)", request.Method, request.URL, request.Body)
	}
	if request.Headers != nil {
		for k, v := range request.Headers {
			req.Header.Set(k, v)
		}
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request (method = %v, url = %v, request body = %v)", request.Method, request.URL, request.Body)
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res, resBody, errors.Wrapf(err, "can not read response (method = %v, url = %v, request body = %v)", request.Method, request.URL, request.Body)
	}
	return res, resBody, nil
}

func NewHTTPClient(timeoutSec int, dnsCacheSec int, idleConnTimeout int, localAddr net.IP) (*HTTPClient) {
	if timeoutSec == 0 {
		timeoutSec = 30
	}
	if dnsCacheSec == 0 {
		dnsCacheSec = 10
	}
	if idleConnTimeout == 0 {
		idleConnTimeout = 180
	}
	newHTTPClient := &HTTPClient{
		timeout:           timeoutSec,
		idleConnTimeout:   idleConnTimeout,
		localAddrIp:       localAddr,
		resolver:          dnscache.New(time.Second * time.Duration(dnsCacheSec)),
		resolverIdx:       0,
		resolverIdxMutex:  new(sync.Mutex),
		clientsCache:      make(map[string]*http.Client),
		clientsCacheMutex: new(sync.Mutex),
	}
	if localAddr != nil {
		newHTTPClient.localAddr = &net.TCPAddr{
			IP: localAddr,
		}
	}
	return newHTTPClient
}


