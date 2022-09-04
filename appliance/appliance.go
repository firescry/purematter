package appliance

import (
	"log"

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
	endpoints, err := client.GenerateEndpoints(host)
	if err != nil {
		log.Fatal(err)
	}

	appliance := Appliance{
		Manufacturer: dev.Device.Manufacturer,
		ModelName:    dev.Device.ModelName,
		ModelNumber:  dev.Device.ModelNumber,
		UUID:         dev.Device.UDN,

		endpoints: endpoints,
	}

	return &appliance
}
