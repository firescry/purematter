package dhe

import (
	"crypto/rand"
	"math/big"
)

type DHE struct {
	g *big.Int // Base
	p *big.Int // Modulus
	a *big.Int // Secret
}

func NewDHE(base, mod, secret *big.Int) *DHE {
	return &DHE{
		g: base,
		p: mod,
		a: secret,
	}
}

func (dh *DHE) GetIntermediate() *big.Int {
	intermediate := &big.Int{}
	intermediate.Exp(dh.g, dh.a, dh.p)
	return intermediate
}

func (dh *DHE) GetSharedKey(intermediate *big.Int) *big.Int {
	key := &big.Int{}
	key.Exp(intermediate, dh.a, dh.p)
	return key
}

func GenDHESecret(bits int) (*big.Int, error) {
	return rand.Prime(rand.Reader, bits)
}
