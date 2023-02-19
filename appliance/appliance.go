package appliance

import (
	"bytes"
	"encoding/base64"
	"log"

	"github.com/firescry/purematter/client"
	"github.com/firescry/purematter/cryptography"

	"github.com/huin/goupnp"
)

type Appliance struct {
	Manufacturer string
	ModelName    string
	ModelNumber  string
	UUID         string

	SecurityApi *client.ApiEndpoint
	AirApi      *client.ApiEndpoint
	FirmwareApi *client.ApiEndpoint
	UserinfoApi *client.ApiEndpoint
	WifiApi     *client.ApiEndpoint
	DeviceApi   *client.ApiEndpoint
	FilterApi   *client.ApiEndpoint
}

func NewAppliance(dev goupnp.RootDevice) *Appliance {
	host := dev.URLBase.Hostname()

	return &Appliance{
		Manufacturer: dev.Device.Manufacturer,
		ModelName:    dev.Device.ModelName,
		ModelNumber:  dev.Device.ModelNumber,
		UUID:         dev.Device.UDN,

		SecurityApi: NewSecurityApi(host),
		AirApi:      NewAirApi(host),
		FirmwareApi: NewFirmwareApi(host),
		UserinfoApi: NewUserinfoApi(host),
		WifiApi:     NewWifiApi(host),
		DeviceApi:   NewDeviceApi(host),
		FilterApi:   NewFilterApi(host),
	}
}

func (a *Appliance) InitConnection() {
	secret, _ := cryptography.GenDHESecret(512)
	dhe := cryptography.NewDHE(
		client.HexToBigInt(PhilipsDHBase),
		client.HexToBigInt(PhilipsDHMod),
		secret)
	localInter := dhe.GetIntermediate()

	request := GetSecurityRequest(localInter)
	body := a.SecurityApi.Put("application/json", bytes.NewReader(request))
	foreingInter, encryptedKey := ParseKeyExResponse(body)

	tmpKeyRaw := dhe.GetSharedKey(foreingInter)
	tmpKey := tmpKeyRaw.Bytes()[:16]

	tmpCrypter := cryptography.NewCrypter(tmpKey)
	key := tmpCrypter.Decrypt(encryptedKey.Bytes())
	key = key[:16]
	crypter := cryptography.NewCrypter(key)

	secEncoded := a.SecurityApi.Get()
	secEncrypted, _ := base64.StdEncoding.DecodeString(string(secEncoded))
	sec := crypter.Decrypt(secEncrypted)
	log.Printf("%s", sec[2:])

	airEncoded := a.AirApi.Get()
	airEncrypted, _ := base64.StdEncoding.DecodeString(string(airEncoded))
	air := crypter.Decrypt(airEncrypted)
	log.Printf("%s", air[2:])

	fwEncoded := a.FirmwareApi.Get()
	fwEncrypted, _ := base64.StdEncoding.DecodeString(string(fwEncoded))
	fw := crypter.Decrypt(fwEncrypted)
	log.Printf("%s", fw[2:])

	uiEncoded := a.UserinfoApi.Get()
	uiEncrypted, _ := base64.StdEncoding.DecodeString(string(uiEncoded))
	ui := crypter.Decrypt(uiEncrypted)
	log.Printf("%s", ui[2:])

	wfEncoded := a.WifiApi.Get()
	wfEncrypted, _ := base64.StdEncoding.DecodeString(string(wfEncoded))
	wf := crypter.Decrypt(wfEncrypted)
	log.Printf("%s", wf[2:])

	dvEncoded := a.DeviceApi.Get()
	dvEncrypted, _ := base64.StdEncoding.DecodeString(string(dvEncoded))
	dv := crypter.Decrypt(dvEncrypted)
	log.Printf("%s", dv[2:])

	frEncoded := a.FilterApi.Get()
	frEncrypted, _ := base64.StdEncoding.DecodeString(string(frEncoded))
	fr := crypter.Decrypt(frEncrypted)
	log.Printf("%s", fr[2:])
}
