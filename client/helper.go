package client

import "math/big"

func HexToBigInt(hex string) *big.Int {
	i := &big.Int{}
	i.SetString(hex, 16)
	return i
}
