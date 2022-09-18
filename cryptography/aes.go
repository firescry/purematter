package cryptography

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

type Crypter struct {
	block cipher.Block
	iv    []byte
}

func NewCrypter(key []byte) *Crypter {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	return &Crypter{
		block: block,
		iv:    make([]byte, aes.BlockSize),
	}
}

func (c *Crypter) Decrypt(data []byte) []byte {
	mode := cipher.NewCBCDecrypter(c.block, c.iv)
	mode.CryptBlocks(data, data)
	return trim(data)
}

// TODO: Handle first two blocks that contain corruperd data due to incorrect iv
func trim(data []byte) []byte {
	term := '\u0004'

	index := bytes.IndexRune(data, term)
	if index > 0 {
		return data[:index]
	}
	return data
}
