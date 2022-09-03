package appliance

import (
	"github.com/huin/goupnp"
)

type Appliance struct {
	host         string
	Manufacturer string
	ModelName    string
	ModelNumber  string
	UUID         string
}

func NewAppliance(dev goupnp.RootDevice) *Appliance {
	return &Appliance{
		host:         dev.URLBase.Hostname(),
		Manufacturer: dev.Device.Manufacturer,
		ModelName:    dev.Device.ModelName,
		ModelNumber:  dev.Device.ModelNumber,
		UUID:         dev.Device.UDN,
	}
}
