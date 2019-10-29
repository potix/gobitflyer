package realtime

import (
        "github.com/potix/gobitflyer/client"
	"github.com/potix/gobitflyer/api/types"
        "github.com/potix/gobitflyer/api/public"
)

type TickerCallback func(productCode types.ProductCode, getTickerResponse *public.GetTickerResponse, callbackData interface{})

type TickerChannel struct {
	ProductCode     types.ProductCode
	WsClient        *client.WSClient
	Callback        TickerCallback
	CallbackData    interface{}
	Subscribed      uint32
	UnsubscribeChan chan *JsonRPC2Subscribe
}

type JsonRPC2TickerNotify struct {
	JsonRpc string                      `json:"jsonrpc"`
	Method  string                      `json:"method"`
	Params  *JsonRPC2TickerNotifyParams `json:"params"`
}

type JsonRPC2TickerNotifyParams struct {
	Channel string                    `json:"channel"`
	Message *public.GetTickerResponse `json:"message"`
}

