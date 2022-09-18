package appliance

import "github.com/firescry/purematter/client"

const (
	WifiApiUrlTemplate = "http://0.0.0.0/di/v1/products/0/wifi"
)

func NewWifiApi(host string) *client.ApiEndpoint {
	return client.NewApiEndpoint(WifiApiUrlTemplate, host)
}
