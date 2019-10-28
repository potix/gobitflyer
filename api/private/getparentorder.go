package private

import (
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
	getParentOrderPath string = "/v1/me/getparentorder"
)

type GetParentOrderResponse  struct {
	Id             int64                      `json:"id"`
	ParentOrderId  string                     `json:"parent_order_id"`
        OrderMethod    types.OrderMethod          `json:"order_method"`
        MinuteToExpire int64                      `json:"minute_to_expire"`
        TimeInForce    types.TimeInForce          `json:"time_in_force"`
        Parameters     []*GetParentOrderParameter `json:"parameters"`
}

type GetParentOrderParameter struct {
        ProductCode    types.ProductCode   `json:"product_code"`
        ConditionType  types.ConditionType `json:"condition_type"`
        Side           types.Side          `json:"side"`
        Price          float64             `json:"price"`
        Size           float64             `json:"size"`
        TriggerPrice   float64             `json:"trigger_price"`
        Offset         float64             `json:"offset"`
}

type GetParentOrderRequest struct {
	Path                    string `url:"-"`
        ParentOrderId           string `url:"parent_order_id,omitempty"`
        ParentOrderAcceptanceId string `url:"parent_order_acceptance_id,omitempty"`
}

func (b *GetParentOrderRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	v, err := query.Values(b)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create query of get parent order request")
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

func NewGetParentOrderRequest(IdType types.IdType, orderId string) (*GetParentOrderRequest, error) {
        switch IdType {
        case types.IdTypeParentOrderId:
                return &GetParentOrderRequest{
                        Path:          getParentOrderPath,
                        ParentOrderId: orderId,
                }, nil
        case types.IdTypeParentOrderAcceptanceId:
                return &GetParentOrderRequest{
                        Path:                    getParentOrderPath,
                        ParentOrderAcceptanceId: orderId,
                }, nil
        default:
                return nil, errors.Errorf("unexpected id type (id type = %v)", IdType)
        }
}
