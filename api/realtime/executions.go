package realtime

import (
        "github.com/potix/gobitflyer/client"
	"github.com/potix/gobitflyer/api/types"
        "github.com/potix/gobitflyer/api/public"
)

type ExecutionsCallback func(productCode types.ProductCode, getExecutionsResponse public.GetExecutionsResponse, callbackData interface{})

type ExecutionsChannel struct {
        ProductCode     types.ProductCode
	WsClient        *client.WSClient
	Callback        ExecutionsCallback
	CallbackData    interface{}
	Subscribed      uint32
	UnsubscribeChan chan *JsonRPC2Subscribe
}

type JsonRPC2ExecutionsNotify struct {
	JsonRpc string                          `json:"jsonrpc"`
	Method  string                          `json:"method"`
	Params  *JsonRPC2ExecutionsNotifyParams `json:"params"`
}

type JsonRPC2ExecutionsNotifyParams struct {
	Channel string                        `json:"channel"`
	Message  public.GetExecutionsResponse `json:"message"`
}

