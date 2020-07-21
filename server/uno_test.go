package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandColor(t *testing.T) {
	testRed := randColor(0)
	assert.Equal(t, testRed, "red") // check for red color

	testBlue := randColor(1)
	assert.Equal(t, testBlue, "blue") // check for blue color

	testGreen := randColor(2)
	assert.Equal(t, testGreen, "green") // check for green color

	testYellow := randColor(3)
	assert.Equal(t, testYellow, "yellow") // check for yellow color

	testEmptyColor := randColor(4)
	assert.Equal(t, testEmptyColor, "") // check for empty color
}

func TestRandomCard(t *testing.T) {

	var colors = []string{"red", "blue", "green", "yellow", ""}

	deck := newRandomCard()
	firstCard := deck[0]

	arrIndex, hasValue := contains(colors, firstCard.Color)

	assert.Equal(t, len(deck), 1) // check that the deck has a card in it
	assert.Equal(t, hasValue, true) // check that the card has a color from colors
	assert.NotEqual(t, arrIndex, -1) // check that the card has a number that is not -1
}
