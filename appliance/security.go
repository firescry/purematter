package appliance

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/firescry/purematter/client"
)

const (
	// Pre-shared Diffie–Hellman parameters
	PhilipsDHBase = "a4d1cbd5c3fd34126765a442efb99905f8104dd258ac507fd6406cff14266d31266fea1e5c41564b777e690f5504f213160217b4b01b886a5e91547f9e2749f4d7fbd7d3b9a92ee1909d0d2263f80a76a6a24c087a091f531dbf0a0169b6a28ad662a4d18e73afa32d779d5918d08bc8858f4dcef97c2a24855e6eeb22b3b2e5"
	PhilipsDHMod  = "b10b8f96a080e01dde92de5eae5d54ec52c99fbcfb06a3c69a6a9dca52d23b616073e28675a23d189838ef1e2ee652c013ecb4aea906112324975c3cd49b83bfaccbdd7d90c4bd7098488e9c219a73724effd6fae5644738faa31a4ff55bccc0a151af5f0dc8b4bd45bf37df365c1a65e68cfda76d4da708df1fb2bc2e4a4371"

	SecurityApiUrlTemplate = "http://0.0.0.0/di/v1/products/0/security"
)

type SecurityRequest struct {
	Diffie string `json:"diffie"`
}

type Security struct {
	Hellman string `json:"hellman"`
	Key     string `json:"key"`
}

func NewSecurityApi(host string) *client.ApiEndpoint {
	return client.NewApiEndpoint(SecurityApiUrlTemplate, host)
}

func GetSecurityRequest(intermediate *big.Int) []byte {
	r := SecurityRequest{
		Diffie: fmt.Sprintf("%x", intermediate),
	}
	j, _ := json.Marshal(r)
	return j
}

func ParseKeyExResponse(data []byte) (*big.Int, *big.Int) {
	r := Security{}
	json.Unmarshal(data, &r)
	inter := client.HexToBigInt(r.Hellman)
	key := client.HexToBigInt(r.Key)
	return inter, key
}
