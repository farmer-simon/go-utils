package egu

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUUID(t *testing.T) {
	fmt.Println(UUID())
}

func TestDefaultString(t *testing.T) {
	defaultV := "someValue"
	assert.Equal(t, defaultV, DefaultString("", defaultV))
	s := "a"
	assert.Equal(t, s, DefaultString(s, ""))
}