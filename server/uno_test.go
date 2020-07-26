package main

import (
	"testing"

	"github.com/jak103/uno/model"
	"github.com/stretchr/testify/assert"
)

// test to make sure the correct colors are being returned from RandColor
// this test is not super helpful - it is being used to get used to unit testing in golang
func TestRandColor(t *testing.T) {
	assert.Equal(t, randColor(0), "red")
	assert.Equal(t, randColor(1), "blue")
	assert.Equal(t, randColor(2), "green")
	assert.Equal(t, randColor(3), "yellow")
	assert.Equal(t, randColor(4), "")
}

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

func TestCheckForWinner(t *testing.T) {
	players = []string{"player1", "player2"}
	allCards = make(map[string][]model.Card)
	dealCards()
	assert.Equal(t, "", checkForWinner())
	allCards[players[0]] = make([]model.Card, 0)
	assert.Equal(t, "player1", checkForWinner())
}
