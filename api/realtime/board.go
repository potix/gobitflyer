package realtime

import (
        "github.com/potix/gobitflyer/client"
        "github.com/potix/gobitflyer/api/types"
        "github.com/potix/gobitflyer/api/public"
)

type BoardCallback func(productCode types.ProductCode, getBoardResponse *public.GetBoardResponse, callbackData interface{})

type BoardChannel struct {
	ProductCode          types.ProductCode
	WsClient             *client.WSClient
	Callback             BoardCallback
	CallbackData         interface{}
	Subscribed           bool
	WriteDataChan        chan *JsonRPC2Subscribe
	Merge                bool
	GetBoardResponseFull *public.GetBoardResponse
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
