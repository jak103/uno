package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	// setup test data
	var data []string = []string{}
	data = append(data, "Test")

	// check to see if "Test" is found in the correct index
	index, found := contains(data, "Test")
	assert.Equal(t, index, 0)
	assert.Equal(t, found, true)

	// make sure a missing string is not found at any index
	index, found = contains(data, "Invalid entry")
	assert.Equal(t, index, -1)
	assert.Equal(t, found, false)
}
