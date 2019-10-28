package private

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
        cancelParentOrderPath string = "/v1/me/cancelparentorder"
)

type CancelParentOrderRequest  struct {
	Path                    string             `json:"-"`
	ProductCode             types.ProductCode  `json:"product_code"`
	ParentOrderId           string             `json:"parent_order_id,omitempty"`
	ParentOrderAcceptanceId string             `json:"parent_order_acceptance_id,omitempty"`
}

func (r *CancelParentOrderRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	body, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create body of cancel parent order request")
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

func NewCancelParentOrderRequest(productCode types.ProductCode, IdType types.IdType, orderId string) (*CancelParentOrderRequest, error) {
	switch IdType {
	case types.IdTypeParentOrderAcceptanceId:
		return &CancelParentOrderRequest{
			Path:                    cancelParentOrderPath,
			ProductCode:             productCode,
			ParentOrderAcceptanceId: orderId,
		}, nil
	case types.IdTypeParentOrderId:
		return &CancelParentOrderRequest{
			Path:          cancelParentOrderPath,
			ProductCode:   productCode,
			ParentOrderId: orderId,
		}, nil
	default:
		return nil, errors.Errorf("unexpected id type (id type = %v)", IdType)
	}
}
