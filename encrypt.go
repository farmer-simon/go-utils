package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/farmer-simon/go-utils/aes"
	"github.com/farmer-simon/go-utils/rsa"
)

//Md5 return the encrypt string by md5 algorithm
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1 return hash string
func Sha1(s string) string {
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(s))
	result := Sha1Inst.Sum([]byte(""))
	return fmt.Sprintf("%x", result)
}

// Encrypt Encrypt
type Encrypt interface {
	// Encrypt 加密
	Encrypt(src []byte) ([]byte, error)

	// Decrypt 解密
	Decrypt(src []byte) ([]byte, error)
}

// Aes encrypt
func Aes(key []byte) Encrypt {
	return aes.New(key)
}

// Rsa encrypt
func Rsa(public, private []byte) Encrypt {
	return rsa.New(public, private)
}

// Base64Decode decode base64 string
func Base64Decode(encoded string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil
	}
	return decoded
}

// Base64Encode return base64 string
func Base64Encode(source []byte) string {
	return base64.StdEncoding.EncodeToString(source)
}
