package utils

import (
	"bytes"
	"crypto/rand"
	uuid "github.com/satori/go.uuid"
	"math/big"
)

// UUID return unique id
func UUID() string {
	return uuid.NewV4().String()
}

// DefaultString return defaultV if v is empty
func DefaultString(v, defaultV string) string {
	if v == "" {
		return defaultV
	}

	return v
}

// RandomString return random string
func RandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
