package public

import (
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/client"
	"github.com/potix/gobitflyer/api/types"
)

const (
        getTickerPath string = "/v1/getticker"
)

type GetTickerResponse struct {
	ProductCode     types.ProductCode `json:"product_code"`
	Timestamp       string            `json:"timestamp"`
	TickId          int64             `json:"tick_id"`
	BestBid         float64           `json:"best_bid"`
	BestAsk         float64           `json:"best_ask"`
	BestBidSize     float64           `json:"best_bid_size"`
	BestAskSize     float64           `json:"best_ask_size"`
	TotalBidDepth   float64           `json:"total_bid_depth"`
	TotalAskDepth   float64           `json:"total_ask_depth"`
	LTP             float64           `json:"ltp"`
	Volume          float64           `json:"volume"`
	VolumeByProduct float64           `json:"volume_by_product"`
}

type GetTickerRequest struct {
	Path        string            `url:"-"`
	ProductCode types.ProductCode `url:"product_code,omitempty"`
}

func (b *GetTickerRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	v, err := query.Values(b)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create query of get ticker request")
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

func NewGetTickerRequest(productCode types.ProductCode) (*GetTickerRequest) {
        return &GetTickerRequest{
                Path:        getTickerPath,
		ProductCode: productCode,
        }
}


