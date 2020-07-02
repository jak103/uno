package main

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

// go test -cover

func TestUno(t *testing.T)  {
	fmt.Println("Running Uno Test...");
	assert := assert.New(t) // for multiple asserts

	a := "Hello"
	b := "Hello"

	assert.Equal(a, b, "The two words should be the same.")
}