package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckForWinner(t *testing.T) {
	players = []string{"Abby", "Bob", "Carol"}
	allCards = make(map[string][]Card)
	dealCards()

	// Assertions
	// making Bob have 0 cards making the checkforwinner func return Bob
	allCards[players[1]] = make([]Card, 0)
	assert.Equal(t, "Bob", checkForWinner())

}

func TestRandColor(t *testing.T) {
	//Assertions
	assert.Equal(t, "red", randColor(0))
	assert.Equal(t, "blue", randColor(1))
	assert.Equal(t, "green", randColor(2))
	assert.Equal(t, "yellow", randColor(3))

	// Test empty color
	assert.Equal(t, "", randColor(-1))
}

func TestCheckID(t *testing.T) {
	// Testing to see if the gameID is empty
	gameID = "Test"
	// Assertions
	assert.NotEqual(t, "", checkID(gameID))
}
