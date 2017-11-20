package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var iv = []byte("bundesliga201711")

func encrypt(text []byte) ([]byte, error) {
	block, _ := aes.NewCipher(secret)
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, len(b))
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext, []byte(b))
	return ciphertext, nil
}

func decrypt(text []byte) ([]byte, error) {
	block, _ := aes.NewCipher(secret)
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
