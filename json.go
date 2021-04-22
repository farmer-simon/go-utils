package egu


import (
	"fmt"
	"github.com/pquerna/ffjson/ffjson"
)

// Encode json stringify
func JsonEncode(v interface{}) (string, error) {
	buf, err := ffjson.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

// Decode json to interface
func JsonDecode(buf []byte, obj interface{}) error {
	return ffjson.Unmarshal(buf, obj)
}

// MustJsonEncode only return string, ignore error
func MustJsonEncode(v interface{}) string  {
	buf, err := ffjson.Marshal(v)
	if err != nil {
		fmt.Printf("json encode %v err: %v", v, err)
	}
	return string(buf)
}