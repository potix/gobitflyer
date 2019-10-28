package public

import (
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
        getExecutionsPath string = "/v1/getexecutions"
)

type GetExecutionsResponse []*GetExecutionsExecution

type GetExecutionsExecution struct {
	Id                         int64   `json:"id"`
	Side                       string  `json:"side"`
	Price                      float64 `json:"price"`
	Size                       float64 `json:"size"`
	ExecDate                   string  `json:"exec_date"`
	BuyChildOrderAcceptanceId  string  `json:"buy_child_order_acceptance_id"`
	SellChildOrderAcceptanceId string  `json:"sell_child_order_acceptance_id"`
}

type GetExecutionsRequest struct {
	Path        string            `url:"-"`
	ProductCode types.ProductCode `url:"product_code,omitempty"`
	types.Pagination
}

func (b *GetExecutionsRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	v, err := query.Values(b)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create query of get executions request")
	}
	query := v.Encode()
	pathQuery := b.Path + "?" + query
        return &client.HTTPRequest {
		PathQuery: pathQuery,
                URL: endpoint + pathQuery,
                Method: "GET",
                Headers: nil,
                Body: nil,
        }, nil
}

func NewGetExecutionsRequest(productCode types.ProductCode, count int64, before int64, after int64) (*GetExecutionsRequest) {
        return &GetExecutionsRequest{
                Path:        getExecutionsPath,
		ProductCode: productCode,
		Pagination:  types.Pagination {
			Count:  count,
			Before: before,
			After:  after,
		},
        }
}


