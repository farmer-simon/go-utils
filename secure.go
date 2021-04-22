package egu


import (
	"fmt"
	"log"
)

// FatalError
func FatalError(msg string, err error) {
	if err != nil {
		panic(fmt.Errorf("%s:%v", msg, err))
	}

	log.Printf("%s Success\n", msg)
}



// SecurePanic only panic when err not nil
func SecurePanic(err error) {
	if err != nil {
		panic(err)
	}
}

