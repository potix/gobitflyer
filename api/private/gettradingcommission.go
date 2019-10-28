package private

import (
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
	getTradingCommissionPath string = "/v1/me/gettradingcommission"
)

type GetTradingCommissionResponse  struct {
	CommissionRate float64 `json:"commission_rate"`
}

type GetTradingCommissionRequest struct {
	Path                   string            `url:"-"`
        ProductCode            types.ProductCode `url:"product_code,omitempty"`
}

func (r *GetTradingCommissionRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	v, err := query.Values(r)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create query of get trading commission request")
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

func NewGetTradingCommissionRequest(productCode types.ProductCode) (*GetTradingCommissionRequest) {
	return &GetTradingCommissionRequest{
		Path:                   getTradingCommissionPath,
		ProductCode:            productCode,
	}
}
