package egu

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCurrentDir(t *testing.T) {
	fmt.Println(GetCurrentDir())
}

func TestGetCurrentPath(t *testing.T) {
	fmt.Printf(GetCurrentPath())
}

func TestPathExist(t *testing.T) {
	assert.True(t, PathExist(GetCurrentPath()))
}

func TestMkdir(t *testing.T) {
	assert.Nil(t, Mkdir(GetCurrentDir(), true))
}

func TestRuntimeCaller(t *testing.T) {
	fmt.Println(RuntimeCaller())
}