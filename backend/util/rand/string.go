package rand

import (
	"encoding/base64"
	"math/rand"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(n int) (string, error) {
	b, err := GenerateRandomBytes(n)

	return base64.URLEncoding.EncodeToString(b), err
}
