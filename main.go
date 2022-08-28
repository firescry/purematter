package main

import (
	"log"
	"net/url"

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
	log.Printf("Found %d device(s):\n", len(maybeRootDevices))
	for _, maybeRootDevice := range maybeRootDevices {
		if maybeRootDevice.Err != nil {
			log.Fatalf("  Failed to probe device at %s\n", maybeRootDevice.Location.String())
		} else {
			locations = append(locations, maybeRootDevice.Location)
			log.Printf("  Successfully probed device at %s\n", maybeRootDevice.Location.String())
		}
	}

	log.Printf("Attempt to re-acquire %d device(s):\n", len(locations))
	for _, location := range locations {
		if rootDevice, err := goupnp.DeviceByURL(location); err != nil {
			log.Fatalf("  Failed to reacquire device at %s: %v\n", location.String(), err)
		} else {
			log.Printf("  Successfully reacquired device at %s:\n", location.String())
			log.Printf("    %s\n", rootDevice.Device.String())
		}
	}

}
