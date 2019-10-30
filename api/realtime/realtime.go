package realtime

import (
        "github.com/potix/gobitflyer/client"
        "github.com/potix/gobitflyer/api/types"
        "github.com/potix/gobitflyer/api/public"
)

type BoardSnapshotCallback func(productCode types.ProductCode, getBoardResponse *public.GetBoardResponse, callbackData interface{})
type BoardCallback func(productCode types.ProductCode, getBoardResponse *public.GetBoardResponse, callbackData interface{})
type TickerCallback func(productCode types.ProductCode, getTickerResponse *public.GetTickerResponse, callbackData interface{})
type ExecutionsCallback func(productCode types.ProductCode, getExecutionsResponse public.GetExecutionsResponse, callbackData interface{})

type RealtimeChannel struct {
	ProductCode           types.ProductCode
	WsClient              *client.WSClient
	RealtimeType          types.RealtimeType
	BoardSnapshotCallback BoardSnapshotCallback
	BoardCallback         BoardCallback
	TickerCallback        TickerCallback
	ExecutionsCallback    ExecutionsCallback
	CallbackData          interface{}
	Subscribed            uint32
	UnsubscribeChan       chan *JsonRPC2Subscribe
	Merge                 bool
	GetBoardResponseFull  *public.GetBoardResponse
}

type JsonRPC2BoardSnapshotNotify struct {
	JsonRpc string                             `json:"jsonrpc"`
	Method  string                             `json:"method"`
	Params  *JsonRPC2BoardSnapshotNotifyParams `json:"params"`
}

type JsonRPC2BoardSnapshotNotifyParams struct {
	Channel string                   `json:"channel"`
	Message *public.GetBoardResponse `json:"message"`
}

type JsonRPC2BoardNotify struct {
	JsonRpc string                     `json:"jsonrpc"`
	Method  string                     `json:"method"`
	Params  *JsonRPC2BoardNotifyParams `json:"params"`
}

type JsonRPC2BoardNotifyParams struct {
	Channel string                   `json:"channel"`
	Message *public.GetBoardResponse `json:"message"`
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

type JsonRPC2ExecutionsNotify struct {
	JsonRpc string                          `json:"jsonrpc"`
	Method  string                          `json:"method"`
	Params  *JsonRPC2ExecutionsNotifyParams `json:"params"`
}

type JsonRPC2ExecutionsNotifyParams struct {
	Channel string                        `json:"channel"`
	Message  public.GetExecutionsResponse `json:"message"`
}

type JsonRPC2Subscribe struct {
	JsonRpc string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Params  JsonRPC2SubscribeParams `json:"params"`
}

type JsonRPC2SubscribeParams struct {
	Channel string `json:"channel"`
}

