package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRandColor(t *testing.T){
	// Testing colors work
	assert.Equal(t, "red", randColor(0))
	assert.Equal(t, "blue", randColor(1))
	assert.Equal(t, "green", randColor(2))
	assert.Equal(t, "yellow", randColor(3))
	assert.Equal(t, "", randColor(4))
}

func TestRandomCard(t *testing.T){
	// Testing 108 random cards
	for i := 0; i < 108; i++ {
		card := newRandomCard()

		// Check that card has color
		assert.NotEqual(t, "", card[0].Color)

		// Check that card has valid value
		assert.Less(t, -1, card[0].Number)
		assert.Greater(t, 10, card[0].Number)
	}
}

func TestContains(t *testing.T){
	array := []string{"Alpha", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot"}

	// Check that an element is found
	index, found := contains(array, "Foxtrot")
	assert.Equal(t, index, 5)
	assert.Equal(t, found, true)

	// Check that an invalid element is not found
	index, found = contains(array, "Invalid")
	assert.Equal(t, index, -1)
	assert.Equal(t, found, false)
}