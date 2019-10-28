package private

import (
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
	getExecutionsPath string = "/v1/me/getexecutions"
)

type GetExecutionsResponse  []*GetExecutionsExecution

type GetExecutionsExecution struct {
	Id                     int64      `json:"id"`
	ChildOrderId           string     `json:"child_order_id"`
	Side                   types.Side `json:"side"`
	Price                  float64    `json:"price"`
	Size                   float64    `json:"size"`
	Commission             float64    `json:"commission"`
	ExecDate               string     `json:"expire_date"`
	ChildOrderAcceptanceId string     `json:"child_order_acceptance_id"`
}

type GetExecutionsRequest struct {
	Path                   string            `url:"-"`
        ProductCode            types.ProductCode `url:"product_code,omitempty"`
	types.Pagination
	ChildOrderId           string            `url:"child_order_id,omitempty"`
	ChildOrderAcceptanceId string            `url:"child_order_acceptance_id,omitempty"`
}

func (r *GetExecutionsRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	v, err := query.Values(r)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create query of get executions request")
	}
	query := v.Encode()
	pathQuery := r.Path + "?" + query
        return &client.HTTPRequest {
		PathQuery: pathQuery,
                URL: endpoint + pathQuery,
                Method: "GET",
                Headers: make(map[string]string),
                Body: nil,
        }, nil
}

func NewGetExecutionsRequest(productCode types.ProductCode, count int64, before int64, after int64) (*GetExecutionsRequest) {
	return &GetExecutionsRequest{
		Path:                   getExecutionsPath,
		ProductCode:            productCode,
		Pagination: types.Pagination {
			Count:  count,
			Before: before,
			After:  after,
		},
	}
}

func NewGetExecutionsRequestById(productCode types.ProductCode, IdType types.IdType, orderId string) (*GetExecutionsRequest, error) {
	switch IdType {
	case types.IdTypeChildOrderId:
		return &GetExecutionsRequest{
			Path:         getExecutionsPath,
			ProductCode:  productCode,
			ChildOrderId: orderId,
		}, nil
	case types.IdTypeChildOrderAcceptanceId:
		return &GetExecutionsRequest{
			Path:                   getExecutionsPath,
			ProductCode:            productCode,
			ChildOrderAcceptanceId: orderId,
		}, nil
	default:
		return nil, errors.Errorf("unexpected id type (id type = %v)", IdType)
	}
}

