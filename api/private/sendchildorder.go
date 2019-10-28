package private

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
        sendChildOrderPath = "/v1/me/sendchildorder"
)

type SendChildOrderResponse struct {
	ChildOrderAcceptanceId string `json:"child_order_acceptance_id"`
}

type SendChildOrderRequest struct {
	Path           string            `json:"-"`
	ProductCode    types.ProductCode `json:"product_code"`
	ChildOrderType types.OrderType   `json:"child_order_type"`
	Side           types.Side        `json:"side"`
	Price          float64           `json:"price,omitempty"`
	Size           float64           `json:"size"`
	MinuteToExpire int64             `json:"minute_to_expire,omitempty"`
	TimeInForce    types.TimeInForce `json:"time_in_force,omitempty"`
}

func (b *SendChildOrderRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	body, err := json.Marshal(b)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create body of send child order request")
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
        return &client.HTTPRequest {
		PathQuery: b.Path,
                URL: endpoint + b.Path,
                Method: "POST",
                Headers: headers,
                Body: body,
        }, nil
}

func NewSendChildOrderRequest(productCode types.ProductCode,
                              childOrderType types.OrderType,
                              side types.Side,
                              price float64,
                              size float64,
                              minuteToExpire int64,
                              timeInForce types.TimeInForce) (*SendChildOrderRequest) {
        return &SendChildOrderRequest{
                Path:           sendChildOrderPath,
		ProductCode:    productCode,
		ChildOrderType: childOrderType,
		Side:           side,
		Price:          price,
		Size:           size,
		MinuteToExpire: minuteToExpire,
		TimeInForce:    timeInForce,
        }
}

