package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

type Crypter struct {
	block cipher.Block
	iv    []byte
}

func NewCrypter(key []byte) Crypter {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	return Crypter{
		block: block,
		iv:    make([]byte, aes.BlockSize),
	}
}

func (c Crypter) Decrypt(data []byte) []byte {
	mode := cipher.NewCBCDecrypter(c.block, c.iv)
	mode.CryptBlocks(data, data)
	data, err := unpad(data)
	if err != nil {
		panic(err)
	}
	return data
}

func unpad(padded []byte) ([]byte, error) {
	length := int(padded[len(padded)-1])
	if length > aes.BlockSize {
		return nil, errors.New("incorrect padding length")
	}
	return padded[:len(padded)-length], nil
}
