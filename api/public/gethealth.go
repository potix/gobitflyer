package public

import (
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
        getHealthPath string = "/v1/gethealth"
)

type GetHealthResponse struct {
	Status string `json:"status"`
}

type GetHealthRequest struct {
	Path        string            `url:"-"`
	ProductCode types.ProductCode `url:"product_code,omitempty"`
}

func (b *GetHealthRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	v, err := query.Values(b)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create query of get health request")
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

func NewGetHealthRequest(productCode types.ProductCode) (*GetHealthRequest) {
        return &GetHealthRequest{
                Path:        getHealthPath,
		ProductCode: productCode,
        }
}
