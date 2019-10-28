package private

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
        cancelAllChildOrdersPath string = "/v1/me/cancelallchildorders"
)

type CancelAllChildOrdersRequest  struct {
	Path                   string             `json:"-"`
	ProductCode            types.ProductCode  `json:"product_code,omitempty"`
}

func (b *CancelAllChildOrdersRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	body, err := json.Marshal(b)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create body of cancel child all orders request")
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

func NewCancelAllChildOrdersRequest(productCode types.ProductCode) (*CancelAllChildOrdersRequest) {
	return &CancelAllChildOrdersRequest{
		Path:        cancelAllChildOrdersPath,
		ProductCode: productCode,
	}
}
