package private

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
        cancelChildOrderPath string = "/v1/me/cancelchildorder"
)

type CancelChildOrderRequest  struct {
	Path                   string             `json:"-"`
	ProductCode            types.ProductCode  `json:"product_code"`
	ChildOrderId           string             `json:"child_order_id,omitempty"`
	ChildOrderAcceptanceId string             `json:"child_order_acceptance_id,omitempty"`
}

func (b *CancelChildOrderRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	body, err := json.Marshal(b)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create body of cancel child order request")
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

func NewCancelChildOrderRequest(productCode types.ProductCode, IdType types.IdType, orderId string) (*CancelChildOrderRequest, error) {
	switch IdType {
	case types.IdTypeChildOrderAcceptanceId:
		return &CancelChildOrderRequest{
			Path:                   cancelChildOrderPath,
			ProductCode:            productCode,
			ChildOrderAcceptanceId: orderId,
		}, nil
	case types.IdTypeChildOrderId:
		return &CancelChildOrderRequest{
			Path:         cancelChildOrderPath,
			ProductCode:  productCode,
			ChildOrderId: orderId,
		}, nil
	default:
		return nil, errors.Errorf("unexpected id type (id type = %v)", IdType)
	}
}
