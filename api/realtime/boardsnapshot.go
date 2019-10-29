package realtime

import (
        "github.com/potix/gobitflyer/client"
        "github.com/potix/gobitflyer/api/types"
        "github.com/potix/gobitflyer/api/public"
)

type BoardSnapshotCallback func(productCode types.ProductCode, getBoardResponse *public.GetBoardResponse, callbackData interface{})

type BoardSnapshotChannel struct {
        ProductCode   types.ProductCode
	WsClient      *client.WSClient
	Callback      BoardSnapshotCallback
	CallbackData  interface{}
	Subscribed    bool
	WriteDataChan chan *JsonRPC2Subscribe
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
