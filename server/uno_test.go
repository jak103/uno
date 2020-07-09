package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRandColor(t *testing.T) {
	// test that the colors work
	assert.Equal(t, "red", randColor(0))
	assert.Equal(t, "blue", randColor(1))
	assert.Equal(t, "green", randColor(2))
	assert.Equal(t, "yellow", randColor(3))

	// test that other input produces an empty string
	assert.Equal(t, "", randColor(-1))
	assert.Equal(t, "", randColor(4))
}

func TestNewRandomCard(t *testing.T) {
	// Try 1000 times
	for i := 0; i < 1000; i++ {
		card := newRandomCard()
		// a valid number should always be produced
		// if our usage of rand.Intn is correct
		assert.Less(t, -1, card.Number)
		assert.Greater(t, 10, card.Number)

		// An actual color should always be produced
		assert.NotEqual(t, "", card.Color)
	}
}

func TestCheckID(t *testing.T) {
	// Do not rely on the bug in uno.go that always sets the gameID to 12234
	gameID = "testingid"
	assert.Equal(t, true, checkID("testingid"))
	assert.Equal(t, false, checkID("wrongid"))
}

func TestContains(t *testing.T) {
	usernames := []string{"pippin", "merry", "eowyn", "faramir"}

	// test that it finds an element
	index, found := contains(usernames, "eowyn")
	assert.Equal(t, 2, index)
	assert.Equal(t, true, found)

	// test that it does not find something that is not an element
	index, found = contains(usernames, "boromir")
	assert.Equal(t, -1, index)
	assert.Equal(t, false, found)
}