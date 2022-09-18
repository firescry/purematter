package appliance

import "github.com/firescry/purematter/client"

const (
	DeviceApiUrlTemplate = "http://0.0.0.0/di/v1/products/1/device"
)

func NewDeviceApi(host string) *client.ApiEndpoint {
	return client.NewApiEndpoint(DeviceApiUrlTemplate, host)
}
