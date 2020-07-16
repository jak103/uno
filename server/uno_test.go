package main

//adding a comment to test my github action

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandColor(t *testing.T) {
	assert.Equal(t, randColor(0), "red")
	assert.Equal(t, randColor(1), "blue")
	assert.Equal(t, randColor(2), "green")
	assert.Equal(t, randColor(3), "yellow")
}
