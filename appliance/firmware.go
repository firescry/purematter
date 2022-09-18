package appliance

import "github.com/firescry/purematter/client"

const (
	FirmwareApiUrlTemplate = "http://0.0.0.0/di/v1/products/0/firmware"
)

func NewFirmwareApi(host string) *client.ApiEndpoint {
	return client.NewApiEndpoint(FirmwareApiUrlTemplate, host)
}
