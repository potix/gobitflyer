package private

import (
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
        getBalancePath string = "/v1/me/getbalance"
)

type GetBalanceResponse []*GetBalanceAsset

type GetBalanceAsset struct {
	CurrencyCode types.CurrencyCode `json:"currency_code"`
	Amount       float64            `json:"amount"`
	Available    float64            `json:"available"`
}

type GetBalanceRequest struct {
	Path string `json:"-"`
}

func (b *GetBalanceRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
        return &client.HTTPRequest {
		PathQuery: b.Path,
                URL: endpoint + b.Path,
                Method: "GET",
                Headers: make(map[string]string),
                Body: nil,
        }, nil
}

func NewGetBalanceRequest() (*GetBalanceRequest) {
        return &GetBalanceRequest{
                Path: getBalancePath,
        }
}


