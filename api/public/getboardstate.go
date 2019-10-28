package public

import (
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
        getBoardStatePath string = "/v1/getboardstate"
)

type GetBoardStateResponse struct {
	Health string             `json:"health"`
	State  string             `json:"state"`
	Data   *GetBoardStateData `json:"data"`
}

type GetBoardStateData struct {
	SpecialQuotation int64 `json:"special_quotation"`
}

type GetBoardStateRequest struct {
	Path        string            `url:"-"`
	ProductCode types.ProductCode `url:"product_code,omitempty"`
}

func (b *GetBoardStateRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	v, err := query.Values(b)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create query of get board state request")
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

func NewGetBoardStateRequest(productCode types.ProductCode) (*GetBoardStateRequest) {
        return &GetBoardStateRequest{
                Path:        getBoardStatePath,
		ProductCode: productCode,
        }
}
