package appliance

import "github.com/firescry/purematter/client"

const (
	AirApiUrlTemplate = "http://0.0.0.0/di/v1/products/1/air"
)

func NewAirApi(host string) *client.ApiEndpoint {
	return client.NewApiEndpoint(AirApiUrlTemplate, host)
}
