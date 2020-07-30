package main

import (
	"testing"

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

// func TestCheckForWinner(t *testing.T) {
// 	players = []string{"player1", "player2"}
// 	allCards = make(map[string][]model.Card)
// 	dealCards()
// 	assert.Equal(t, "", checkForWinner())
// 	allCards[players[0]] = make([]model.Card, 0)
// 	assert.Equal(t, "player1", checkForWinner())
// }
