package private

import (
	"github.com/potix/gobitflyer/client"
)

const (
        getPermissionsPath = "/v1/me/getpermissions"
)

type GetPermissionsResponse []string

type GetPermissionsRequest struct {
	Path string `json:"-"`
}

func (b *GetPermissionsRequest) CreateHTTPRequest(endpoint string) (*client.HTTPRequest, error) {
        return &client.HTTPRequest {
		PathQuery: b.Path,
                URL: endpoint + b.Path,
                Method: "GET",
                Headers: make(map[string]string),
                Body: nil,
        }, nil
}

func NewGetPermissionsRequest() (*GetPermissionsRequest) {
        return &GetPermissionsRequest{
                Path: getPermissionsPath,
        }
}


