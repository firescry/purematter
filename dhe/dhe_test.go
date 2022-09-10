package dhe

import (
	"math/big"
	"testing"
)

func TestIntermediateCalculation(t *testing.T) {
	base := big.NewInt(5)
	mod := big.NewInt(23)
	secret := big.NewInt(3)

	dhe := NewDHE(base, mod, secret)

	result := dhe.GetIntermediate()
	expected := big.NewInt(10)
	if expected.Cmp(result) != 0 {
		t.Errorf("Expected value is %d, but got %d!\n", expected, result)
	}
}

func TestSharedKeyCalculation(t *testing.T) {
	base := big.NewInt(5)
	mod := big.NewInt(23)
	secret := big.NewInt(3)
	intermediate := big.NewInt(4)

	dhe := NewDHE(base, mod, secret)

	result := dhe.GetSharedKey(intermediate)
	expected := big.NewInt(18)
	if expected.Cmp(result) != 0 {
		t.Errorf("Expected value is %d, but got %d!\n", expected, result)
	}
}

func TestSecretGeneration(t *testing.T) {
	bits_exp := 32
	secret, err := GenDHESecret(bits_exp)

	if err != nil {
		t.Errorf("Method returned an error for %d bit(s): %s\n", bits_exp, err)
	}

	bits := secret.BitLen()
	if bits_exp != bits {
		t.Errorf("Bit length of secret (%d) is shorter than expected (%d)!\n", bits, bits_exp)
	}
}
