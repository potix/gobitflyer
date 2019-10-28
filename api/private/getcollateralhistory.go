package private

import (
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
	getCollateralHistoryPath string = "/v1/me/getcollateralhistory"
)

type GetCollateralHistoryResponse  []*GetCollateralHistoryEvent

type GetCollateralHistoryEvent struct {
	Id           int64              `json:"id"`
	CurrencyCode types.CurrencyCode `json:"currency_code"`
	Change       float64            `json:"change"`
	Amount       float64            `json:"amount"`
	ReasonCode   string             `json:"reason_code"`
	Date         string             `json:"date"`
}

type GetCollateralHistoryRequest struct {
	Path             string            `url:"-"`
	types.Pagination
}

func (r *GetCollateralHistoryRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	v, err := query.Values(r)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create query of get collateral history")
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

func NewGetCollateralHistoryRequest(count int64, before int64, after int64) (*GetCollateralHistoryRequest) {
	return &GetCollateralHistoryRequest{
		Path:        getCollateralHistoryPath,
		Pagination:  types.Pagination {
			Count:  count,
			Before: before,
			After:  after,
		},
	}
}

