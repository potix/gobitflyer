package private

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
        sendParentOrderPath = "/v1/me/sendparentorder"
)

type SendParentOrderResponse struct {
	ParentOrderAcceptanceId string `json:"parent_order_acceptance_id"`
}

type SendParentOrderRequest struct {
	Path           string                      `json:"-"`
	OrderMethod    types.OrderMethod           `json:"order_method,omitempty"`
	MinuteToExpire int64                       `json:"minute_to_expire,omitempty"`
	TimeInForce    types.TimeInForce           `json:"time_in_force,omitempty"`
	Parameters     []*SendParentOrderParameter `json:"parameters"`
}

type SendParentOrderParameter struct {
	ProductCode    types.ProductCode   `json:"product_code"`
	ConditionType  types.ConditionType `json:"condition_type"`
	Side           types.Side          `json:"side"`
	Price          float64             `json:"price,omitempty"`
	Size           float64             `json:"size"`
	TriggerPrice   float64             `json:"trigger_price,omitempty"`
	Offset         float64             `json:"offset,omitempty"`
}

func (r *SendParentOrderRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	body, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create body of send paent order request")
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
        return &client.HTTPRequest {
		PathQuery: r.Path,
                URL: endpoint + r.Path,
                Method: "POST",
                Headers: headers,
                Body: body,
        }, nil
}

func NewSendParentOrderRequest(orderMethod types.OrderMethod, minuteToExpire int64, timeInForce types.TimeInForce, parameters ...*SendParentOrderParameter) (*SendParentOrderRequest) {
        return &SendParentOrderRequest{
                Path:           sendParentOrderPath,
		OrderMethod:    orderMethod,
		MinuteToExpire: minuteToExpire,
		TimeInForce:    timeInForce,
		Parameters:     parameters,
        }
}

