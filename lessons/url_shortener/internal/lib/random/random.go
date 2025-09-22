package random

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func randomStringUrl(length int) (string, error) {
	if length < 1 {
		return "", fmt.Errorf("length too small")
	}
	if length > 255 {
		return "", fmt.Errorf("length too large")
	}

	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b)[:length], nil
}
