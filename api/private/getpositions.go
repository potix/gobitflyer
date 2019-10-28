package private

import (
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
	getPositionsPath string = "/v1/me/getpositions"
)

type GetPositionsResponse  []*GetPositionsPosition

type GetPositionsPosition struct {
	ProductCode         types.ProductCode `json:"product_code"`
	Side                types.Side        `json:"side"`
	Price               float64           `json:"price"`
	Size                float64           `json:"size"`
	Commission          float64           `json:"commission"`
	SwapPointAccumulate float64           `json:"swap_point_accumulate"`
        RequireCollateral   float64           `json:"require_collateral"`
        OpenDate            string            `json:"open_date"`
        Leverage            float64           `json:"leverage"`
        Pnl                 float64           `json:"pnl"`
        Sfd                 float64           `json:"sfd"`
}

type GetPositionsRequest struct {
	Path                   string            `url:"-"`
        ProductCode            types.ProductCode `url:"product_code"`
}

func (r *GetPositionsRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	v, err := query.Values(r)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create query of get positions")
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

func NewGetPositionsRequest() (*GetPositionsRequest) {
	return &GetPositionsRequest{
		Path:        getPositionsPath,
		ProductCode: "FX_BTC_JPY",
	}
}

