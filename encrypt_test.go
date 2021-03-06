package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	fmt.Println(Md5("test"))
	fmt.Println(Sha1("test"))
}

func TestBase64(t *testing.T) {
	source := []byte("test")
	encode := Base64Encode(source)
	fmt.Println(encode)
	decode := Base64Decode(encode)
	fmt.Println(string(decode))
	assert.Equal(t, source, decode)
}
