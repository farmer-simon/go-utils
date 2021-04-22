package egu

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	assert.Equal(t, 2, Max(1, 2))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min(1, 2))
}

func TestDefault(t *testing.T) {
	assert.Equal(t, 1, DefaultInt(1, 2))
	assert.Equal(t, 1, DefaultInt(0, 1))
}

func TestRound(t *testing.T) {
	assert.Equal(t, 10, Round(10.23))
}

func TestDiv(t *testing.T) {
	assert.Equal(t, 4, Div(12, 3))
}

func TestRandInt(t *testing.T) {
	fmt.Println(RandInt(1, 3))
}