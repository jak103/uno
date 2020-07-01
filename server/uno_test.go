package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandColor(t *testing.T) {
	fmt.Println("Testing Random Color:")
	assert.Equal(t, "red", randColor(0))
	assert.Equal(t, "blue", randColor(1))
	assert.Equal(t, "green", randColor(2))
	assert.Equal(t, "yellow", randColor(3))
	assert.Equal(t, "", randColor(-1))
	assert.Equal(t, "", randColor(4))
}

func TestContains(t *testing.T) {
	fmt.Println("Testing Contains:")
	list := []string{"A", "B", "C", "D"}
	idx, found := contains(list, "A")
	assert.True(t, found)
	assert.Equal(t, 0, idx)
}
