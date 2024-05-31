package helper

import (
	"crypto/rand"
	"encoding/base64"
)

func GeneratedRandomToken(n int) string {
	token := make([]byte, n)
	_, err := rand.Read(token)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(token)
}
