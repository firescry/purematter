package appliance

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type Security struct {
	Diffie  string `json:"diffie,omitempty"`
	Hellman string `json:"hellman,omitempty"`
	Key     string `json:"key,omitempty"`
}

func GetSecurityRequest(intermediate *big.Int) []byte {
	r := Security{
		Diffie: fmt.Sprintf("%x", intermediate),
	}
	j, _ := json.Marshal(r)
	return j
}

func ParseKeyExResponse(data []byte) (*big.Int, *big.Int) {
	r := Security{}
	err := json.Unmarshal(data, &r)
	if err != nil {
		panic(err)
	}
	inter, _ := new(big.Int).SetString(r.Hellman, 16)
	key, _ := new(big.Int).SetString(r.Key, 16)
	return inter, key
}
