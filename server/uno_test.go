package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandColor(t *testing.T) {
	colors := [4]string{"red", "blue", "green", "yellow"}
	for c := 0; c < len(colors); c++ {
		assert.Equal(t, colors[c], randColor(c))
	}
	assert.Equal(t, "", randColor(5))
}

func TestContains(t *testing.T) {
	arr := []string{"1", "2"}
	i, found := contains(arr, "1")
	assert.Equal(t, 0, i)
	assert.Equal(t, true, found)
	i, found = contains(arr, "5")
	assert.Equal(t, -1, i)
	assert.Equal(t, false, found)
}