package private

import (
	"github.com/potix/gobitflyer/client"
)

const (
        getCollateralPath string = "/v1/me/getcollateral"
)

type GetCollateralResponse struct {
	Collateral        float64 `json:"collateral"`
	OpenPositionPNL   float64 `json:"open_position_pnl"`
	RequireCollateral float64 `json:"require_collateral"`
	KeepRate          float64 `json:"keep_rate"`
}

type GetCollateralRequest struct {
	Path string `json:"-"`
}

func (b *GetCollateralRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
        return &client.HTTPRequest {
		PathQuery: b.Path,
                URL: endpoint + b.Path,
                Method: "GET",
                Headers: make(map[string]string),
                Body: nil,
        }, nil
}

func NewGetCollateralRequest() (*GetCollateralRequest) {
        return &GetCollateralRequest{
                Path: getCollateralPath,
        }
}


