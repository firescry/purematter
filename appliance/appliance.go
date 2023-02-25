package appliance

import (
	"bytes"
	"encoding/base64"
	"log"

	"github.com/firescry/purematter/client"
	"github.com/firescry/purematter/cryptography"
)

var epTemplates = map[string]string{
	"air":      "http://0.0.0.0/di/v1/products/1/air",
	"device":   "http://0.0.0.0/di/v1/products/1/device",
	"fltsts":   "http://0.0.0.0/di/v1/products/1/fltsts",
	"firmware": "http://0.0.0.0/di/v1/products/0/firmware",
	"security": "http://0.0.0.0/di/v1/products/0/security",
	"userinfo": "http://0.0.0.0/di/v1/products/0/userinfo",
	"wifi":     "http://0.0.0.0/di/v1/products/0/wifi",
}

type Appliance struct {
	ep map[string]client.ApiEndpoint
}

func NewAppliance() *Appliance {
	return new(Appliance)
}

func (a *Appliance) SetHost(host string) {
	if a.ep == nil {
		a.ep = make(map[string]client.ApiEndpoint)
	}

	for name, template := range epTemplates {
		a.ep[name] = client.NewApiEndpoint(template, host)
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
	body := a.ep["security"].Put("application/json", bytes.NewReader(request))
	foreingInter, encryptedKey := ParseKeyExResponse(body)

	tmpKeyRaw := dhe.GetSharedKey(foreingInter)
	tmpKey := tmpKeyRaw.Bytes()[:16]

	tmpCrypter := cryptography.NewCrypter(tmpKey)
	key := tmpCrypter.Decrypt(encryptedKey.Bytes())
	key = key[:16]
	crypter := cryptography.NewCrypter(key)

	secEncoded := a.ep["security"].Get()
	secEncrypted, _ := base64.StdEncoding.DecodeString(string(secEncoded))
	sec := crypter.Decrypt(secEncrypted)
	log.Printf("%s", sec[2:])

	airEncoded := a.ep["air"].Get()
	airEncrypted, _ := base64.StdEncoding.DecodeString(string(airEncoded))
	air := crypter.Decrypt(airEncrypted)
	log.Printf("%s", air[2:])

	fwEncoded := a.ep["firmware"].Get()
	fwEncrypted, _ := base64.StdEncoding.DecodeString(string(fwEncoded))
	fw := crypter.Decrypt(fwEncrypted)
	log.Printf("%s", fw[2:])

	uiEncoded := a.ep["userinfo"].Get()
	uiEncrypted, _ := base64.StdEncoding.DecodeString(string(uiEncoded))
	ui := crypter.Decrypt(uiEncrypted)
	log.Printf("%s", ui[2:])

	wfEncoded := a.ep["wifi"].Get()
	wfEncrypted, _ := base64.StdEncoding.DecodeString(string(wfEncoded))
	wf := crypter.Decrypt(wfEncrypted)
	log.Printf("%s", wf[2:])

	dvEncoded := a.ep["device"].Get()
	dvEncrypted, _ := base64.StdEncoding.DecodeString(string(dvEncoded))
	dv := crypter.Decrypt(dvEncrypted)
	log.Printf("%s", dv[2:])

	frEncoded := a.ep["fltsts"].Get()
	frEncrypted, _ := base64.StdEncoding.DecodeString(string(frEncoded))
	fr := crypter.Decrypt(frEncrypted)
	log.Printf("%s", fr[2:])
}
