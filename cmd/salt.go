package cmd

import (
	"crypto/sha512"
)

var saltLength = 25

func salt(text []byte, secret []byte) ([]byte) {
	buf := []byte{}
	buf = append(buf, secret...)
	buf = append(buf, text...)
	checksum := sha512.Sum512(buf)
	return checksum[0:saltLength]
}
