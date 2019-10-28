package private

import (
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
	getParentOrdersPath string = "/v1/me/getparentorders"
)

type GetParentOrdersResponse  []*GetParentOrdersOrder

type GetParentOrdersOrder struct {
	Id                      int64             `json:"id"`
	ParentOrderId           string            `json:"parent_order_id"`
	ProductCode             types.ProductCode `json:"product_code"`
	Side                    string            `json:"side"`
	ParentOrderType         string            `json:"parent_order_type"`
	Price                   float64           `json:"price"`
	AveragePrice            float64           `json:"average_price"`
	Size                    float64           `json:"size"`
	ParentOrderState        string            `json:"parent_order_state"`
	ExpireDate              string            `json:"expire_date"`
	ParentOrderDate         string            `json:"parent_order_date"`
	ParentOrderAcceptanceId string            `json:"parent_order_acceptance_id"`
	OutstandingSize         float64           `json:"outstanding_size"`
	CancelSize              float64           `json:"cancel_size"`
	ExecutedSize            float64           `json:"executed_size"`
	TotalCommission         float64           `json:"total_commission"`
}

type GetParentOrdersRequest struct {
	Path              string                 `url:"-"`
        ProductCode       types.ProductCode      `url:"product_code,omitempty"`
	types.Pagination
	ParentOrderState  types.OrderState       `url:"parent_order_state,omitempty"`
}

func (b *GetParentOrdersRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	v, err := query.Values(b)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create query of get parent orders request")
	}
	query := v.Encode()
	pathQuery := b.Path + "?" + query
        return &client.HTTPRequest {
		PathQuery: pathQuery,
                URL: endpoint + pathQuery,
                Method: "GET",
                Headers: make(map[string]string),
                Body: nil,
        }, nil
}

func NewGetParentOrdersRequest(productCode types.ProductCode, count int64, before int64, after int64, orderState types.OrderState) (*GetParentOrdersRequest) {
	return &GetParentOrdersRequest{
		Path:            getParentOrdersPath,
		ProductCode:     productCode,
		Pagination:      types.Pagination {
			Count:  count,
			Before: before,
			After:  after,
		},
	        ParentOrderState: orderState,
	}
}
