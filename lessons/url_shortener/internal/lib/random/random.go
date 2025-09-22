package random

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_")

func RandomStringUrl(length int) (string, error) {
	if length < 1 {
		return "", fmt.Errorf("length too small")
	}
	if length > 255 {
		return "", fmt.Errorf("length too large")
	}
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b), nil
}
