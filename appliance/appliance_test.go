package appliance

import (
	"testing"

	"github.com/huin/goupnp"
)

func TestNewAppliance(t *testing.T) {
	manu := "TestManufacturer"
	name := "TestName"
	numb := "TestNumber"
	udn := "TestUDN"

	dev := goupnp.Device{
		Manufacturer: manu,
		ModelName:    name,
		ModelNumber:  numb,
		UDN:          udn,
	}
	rootDev := goupnp.RootDevice{
		Device: dev,
	}
	app := NewAppliance(rootDev)

	if manu != app.Manufacturer {
		t.Errorf("Expected '%s', got '%s'\n", manu, app.Manufacturer)
	}
	if name != app.ModelName {
		t.Errorf("Expected '%s', got '%s'\n", name, app.ModelName)
	}
	if numb != app.ModelNumber {
		t.Errorf("Expected '%s', got '%s'\n", numb, app.ModelNumber)
	}
	if udn != app.UUID {
		t.Errorf("Expected '%s', got '%s'\n", udn, app.UUID)
	}
}
