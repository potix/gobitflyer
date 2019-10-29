package realtime

type JsonRPC2Subscribe struct {
	JsonRpc string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Params  JsonRPC2SubscribeParams `json:"params"`
}

type JsonRPC2SubscribeParams struct {
	Channel string `json:"channel"`
}
