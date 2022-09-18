package appliance

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetSecurityRequest(t *testing.T) {
	intermediate := big.NewInt(12345678901234)
	result := GetSecurityRequest(intermediate)
	expected := []byte("{\"diffie\":\"b3a73ce2ff2\"}")
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected value is '%s', but got '%s'!\n", expected, result)
	}
}

func TestParseKeyExResponse(t *testing.T) {
	response := []byte("{\"hellman\":\"1234567890abcdef\",\"key\":\"123abc\"}")
	inter, key := ParseKeyExResponse(response)
	expInter := big.NewInt(1311768467294899695)
	expKey := big.NewInt(1194684)
	if expInter.Cmp(inter) != 0 {
		t.Errorf("Expected intermiediate %d, but got %d!\n", expInter, inter)
	}
	if expKey.Cmp(key) != 0 {
		t.Errorf("Expected key %d, but got %d!\n", expKey, key)
	}
}
