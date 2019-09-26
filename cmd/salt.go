package cmd

import (
	"crypto/rand"
)

var saltLength = 16

func salt() ([]byte, error) {
	b := make([]byte, saltLength)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}
