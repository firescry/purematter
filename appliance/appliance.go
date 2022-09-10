package appliance

import (
	"github.com/firescry/purematter/client"

	"github.com/huin/goupnp"
)

type Appliance struct {
	Manufacturer string
	ModelName    string
	ModelNumber  string
	UUID         string

	endpoints map[string]string
}

func NewAppliance(dev goupnp.RootDevice) *Appliance {
	host := dev.URLBase.Hostname()
	endpoints, _ := client.GenerateEndpoints(host)

	return &Appliance{
		Manufacturer: dev.Device.Manufacturer,
		ModelName:    dev.Device.ModelName,
		ModelNumber:  dev.Device.ModelNumber,
		UUID:         dev.Device.UDN,

		endpoints: endpoints,
	}
}
