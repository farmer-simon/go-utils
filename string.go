package utils

import (
	"bytes"
	"crypto/rand"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
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

//GbkToUtf8 Convert GBK to utf8 encoding
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//Utf8ToGbk Convert utf8 to GBK encoding
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//Substring Intercept sub string
func Substring(s string, start, length int) string {
	rs := []rune(s)
	return string(rs[start:length])
}
