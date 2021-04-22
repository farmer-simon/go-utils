package egu


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Obj struct {
	Name string `json:"name"`
}

func TestJson(t *testing.T) {
	o := Obj{
		Name: "test",
	}

	s, err := JsonEncode(o)
	assert.Nil(t, err)
	assert.Equal(t, s, MustJsonEncode(o))

	var d Obj
	assert.Nil(t, JsonDecode([]byte(s), &d))
	assert.Equal(t, o, d)



}

