package private

import (
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/client"
)

const (
	getBalanceHistoryPath string = "/v1/me/getbalancehistory"
)

type GetBalanceHistoryResponse  []*GetBalanceHistoryEvent

type GetBalanceHistoryEvent struct {
	Id           int64              `json:"id"`
	TradeDate    string             `json:"trade_date"`
	ProductCode  types.ProductCode  `json:"product_code"`
	CurrencyCode types.CurrencyCode `json:"currency_code"`
	TradeType    types.TradeType    `json:"trade_type"`
	Price        float64            `json:"price"`
	Amount       float64            `json:"amount"`
	Quantity     float64            `json:"quantity"`
	Commission   float64            `json:"commission"`
	Balance      float64            `json:"balance"`
	OrderId      string             `json:"order_id"`
}

type GetBalanceHistoryRequest struct {
	Path             string             `url:"-"`
        CurrencyCode     types.CurrencyCode `url:"currency_code"`
	types.Pagination
}

func (r *GetBalanceHistoryRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	v, err := query.Values(r)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create query of get balance history")
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

func NewGetBalanceHistoryRequest(currencyCode types.CurrencyCode, count int64, before int64, after int64) (*GetBalanceHistoryRequest) {
	return &GetBalanceHistoryRequest{
		Path:         getBalanceHistoryPath,
		CurrencyCode: currencyCode,
		Pagination:   types.Pagination {
			Count:  count,
			Before: before,
			After:  after,
		},
	}
}

