package public

import (
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
        getBoardPath string = "/v1/getboard"
)

type GetBoardResponse struct {
	MidPrice float64         `json:"mid_price"`
	Bids     []*GetBoardBook `json:"bids"`
	Asks     []*GetBoardBook `json:"asks"`
}

type GetBoardBook struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

type GetBoardRequest struct {
	Path        string            `url:"-"`
	ProductCode types.ProductCode `url:"product_code,omitempty"`
}

func (b *GetBoardRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
        v, err := query.Values(b)
        if err != nil {
                return nil, errors.Wrapf(err, "can not create query of get board request")
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

func NewGetBoardRequest(productCode types.ProductCode) (*GetBoardRequest) {
        return &GetBoardRequest{
                Path:        getBoardPath,
		ProductCode: productCode,
        }
}


