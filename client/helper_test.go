package client

import (
	"math/big"
	"testing"
)

func TestHexToBigInt(t *testing.T) {
	str := "123abc"
	exp := big.NewInt(1194684)
	res := HexToBigInt(str)
	if exp.Cmp(res) != 0 {
		t.Errorf("Expected value %d, but got %d!\n", exp, res)
	}
}
