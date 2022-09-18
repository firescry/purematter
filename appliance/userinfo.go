package appliance

import "github.com/firescry/purematter/client"

const (
	UserinfoApiUrlTemplate = "http://0.0.0.0/di/v1/products/0/userinfo"
)

func NewUserinfoApi(host string) *client.ApiEndpoint {
	return client.NewApiEndpoint(UserinfoApiUrlTemplate, host)
}
