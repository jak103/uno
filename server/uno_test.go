package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -cover

func TestColor(t *testing.T) {
	assert := assert.New(t)

	r1 := randColor(0)
	r2 := ""

	// expected, actual, msg
	assert.NotEqual(r2, r1, "The two colors should be the same.")
}
