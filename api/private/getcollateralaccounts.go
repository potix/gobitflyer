package private

import (
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
	getCollateralAccountsPath string = "/v1/me/getcollateralaccounts"
)

type GetCollateralAccountsResponse []*GetCollateralAccountsAccount

type GetCollateralAccountsAccount struct {
	CurrencyCode types.CurrencyCode `json:"currency_code"`
	Amount       float64            `json:"amount"`
}

type GetCollateralAccountsRequest struct {
	Path string `json:"-"`
}

func (b *GetCollateralAccountsRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
        return &client.HTTPRequest {
		PathQuery: b.Path,
                URL: endpoint + b.Path,
                Method: "GET",
                Headers: make(map[string]string),
                Body: nil,
        }, nil
}

func NewGetCollateralAccountsRequest() (*GetCollateralAccountsRequest) {
        return &GetCollateralAccountsRequest{
                Path: getCollateralAccountsPath,
        }
}


