package public

import (
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/client"
)

const (
        getChatsPath string = "/v1/getchats"
)

type GetChatsResponse []*GetChatsChat

type GetChatsChat struct {
	Nickname string `json:"nickname"`
	Message  string `json:"message"`
	Date     string `json:"date"`
}

type GetChatsRequest struct {
	Path     string `url:"-"`
	FromDate int64  `url:"from_date,omitempty"`
}

func (r *GetChatsRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
	v, err := query.Values(r)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create query of get chats request")
	}
	query := v.Encode()
	pathQuery := r.Path + "?" + query
        return &client.HTTPRequest {
		PathQuery: pathQuery,
                URL: endpoint + pathQuery,
                Method: "GET",
                Headers: nil,
                Body: nil,
        }, nil
}

func NewGetChatsRequest(fromDate int64) (*GetChatsRequest) {
        return &GetChatsRequest{
                Path:     getChatsPath,
		FromDate: fromDate,
        }
}


