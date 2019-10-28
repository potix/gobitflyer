package private

import (
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
	getChildOrdersPath string = "/v1/me/getchildorders"
)

type GetChildOrdersResponse  []*GetChildOrdersOrder

type GetChildOrdersOrder struct {
	Id                     int64             `json:"id"`
	ChildOrderId           string            `json:"child_order_id"`
	ProductCode            types.ProductCode `json:"product_code"`
	Side                   string            `json:"side"`
	ChildOrderType         string            `json:"child_order_type"`
	Price                  float64           `json:"price"`
	AveragePrice           float64           `json:"average_price"`
	Size                   float64           `json:"size"`
	ChildOrderState        string            `json:"child_order_state"`
	ExpireDate             string            `json:"expire_date"`
	ChildOrderDate         string            `json:"child_order_date"`
	ChildOrderAcceptanceId string            `json:"child_order_acceptance_id"`
	OutstandingSize        float64           `json:"outstanding_size"`
	CancelSize             float64           `json:"cancel_size"`
	ExecutedSize           float64           `json:"executed_size"`
	TotalCommission        float64           `json:"total_commission"`
}

type GetChildOrdersRequest struct {
	Path                   string                 `url:"-"`
        ProductCode            types.ProductCode      `url:"product_code,omitempty"`
	types.Pagination
	ChildOrderState        types.OrderState       `url:"child_order_state,omitempty"`
	ChildOrderId           string                 `url:"child_order_id,omitempty"`
	ChildOrderAcceptanceId string                 `url:"child_order_acceptance_id,omitempty"`
	ParentOrderId          string                 `url:"parent_order_id,omitempty"`
}

func (b *GetChildOrdersRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	v, err := query.Values(b)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create query of get child orders request")
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

func NewGetChildOrdersRequest(productCode types.ProductCode, count int64, before int64, after int64, orderState types.OrderState) (*GetChildOrdersRequest) {
	return &GetChildOrdersRequest{
		Path:            getChildOrdersPath,
		ProductCode:     productCode,
		Pagination:      types.Pagination {
			Count:  count,
			Before: before,
			After:  after,
		},
	        ChildOrderState: orderState,
	}
}

func NewGetChildOrdersRequestById(productCode types.ProductCode, IdType types.IdType, orderId string) (*GetChildOrdersRequest, error) {
	switch IdType {
	case types.IdTypeChildOrderId:
		return &GetChildOrdersRequest{
			Path:         getChildOrdersPath,
			ProductCode:  productCode,
			ChildOrderId: orderId,
		}, nil
	case types.IdTypeChildOrderAcceptanceId:
		return &GetChildOrdersRequest{
			Path:                   getChildOrdersPath,
			ProductCode:            productCode,
			ChildOrderAcceptanceId: orderId,
		}, nil
	case types.IdTypeParentOrderId:
		return &GetChildOrdersRequest{
			Path:          getChildOrdersPath,
			ProductCode:   productCode,
			ParentOrderId: orderId,
		}, nil
	default:
		return nil, errors.Errorf("unexpected id type (id type = %v)", IdType)
	}
}

