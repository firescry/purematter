package main

import (
	"log"
	"net/url"

	"github.com/firescry/purematter/appliance"

	"github.com/huin/goupnp"
)

const (
	URN_Philips_DiProduct_1 = "urn:philips-com:device:DiProduct:1"
)

func main() {

	log.Printf("Looking for known devices...\n")
	maybeRootDevices, err := goupnp.DiscoverDevices(URN_Philips_DiProduct_1)
	if err != nil {
		log.Fatalf("Could not discover %s devices: %v\n", URN_Philips_DiProduct_1, err)
	}

	locations := make([]*url.URL, 0, len(maybeRootDevices))
	for _, maybeRootDevice := range maybeRootDevices {
		if maybeRootDevice.Err != nil {
			log.Fatalf("  Failed to probe device at %s\n", maybeRootDevice.Location.String())
		} else {
			locations = append(locations, maybeRootDevice.Location)
		}
	}

	appliances := make([]*appliance.Appliance, 0, len(maybeRootDevices))
	for _, location := range locations {
		if rootDevice, err := goupnp.DeviceByURL(location); err != nil {
			log.Fatalf("Failed to reacquire device at %s: %v\n", location.String(), err)
		} else {
			appliances = append(appliances, appliance.NewAppliance(*rootDevice))
		}
	}

	log.Printf("Found %d device(s):\n", len(appliances))
	for _, appliance := range appliances {
		log.Printf("  Manufacturer: %s\n", appliance.Manufacturer)
		log.Printf("  Name: %s\n", appliance.ModelName)
		log.Printf("  Number: %s\n", appliance.ModelNumber)
		log.Printf("---\n")
	}
}
