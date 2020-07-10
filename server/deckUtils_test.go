package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenerateShuffledDeck(t *testing.T) {
	deck := generateShuffledDeck()

	// check that the deck has the right number of each color
	colorCounts := map[string]int {
		"red": 0,
		"blue": 0,
		"green": 0,
		"yellow": 0,
		"wild": 0,
	}

	for _, card := range deck {
		colorCounts[card.Color]++
	}
	
	assert.Equal(t, 25, colorCounts["red"])
	assert.Equal(t, 25, colorCounts["blue"])
	assert.Equal(t, 25, colorCounts["green"])
	assert.Equal(t, 25, colorCounts["yellow"])
	assert.Equal(t, 8, colorCounts["wild"])

	// check that the deck has the right number of total cards
	assert.Equal(t, 108, len(deck))
}