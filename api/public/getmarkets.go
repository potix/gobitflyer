package public

import (
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
        getMarketsPath string = "/v1/getmarkets"
)

type GetMarketsResponse []*GetMarketsMarket

type GetMarketsMarket struct {
	ProductCode types.ProductCode `json:"product_code"`
	Alias       types.ProductCode `json:"alias"`
}

type GetMarketsRequest struct{
	Path string `url:"-"`
}

func (m *GetMarketsRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	return &client.HTTPRequest {
		URL: endpoint + m.Path,
		Method: "GET",
		Headers: nil,
		Body: nil,
	}, nil
}

func NewGetMarketsRequest() (*GetMarketsRequest) {
	return &GetMarketsRequest{
		Path: getMarketsPath,
	}
}
