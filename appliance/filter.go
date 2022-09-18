package appliance

import "github.com/firescry/purematter/client"

const (
	FilterApiUrlTemplate = "http://0.0.0.0/di/v1/products/1/fltsts"
)

func NewFilterApi(host string) *client.ApiEndpoint {
	return client.NewApiEndpoint(FilterApiUrlTemplate, host)
}
