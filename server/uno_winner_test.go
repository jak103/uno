package main

  import (
 	"testing"

  	"github.com/stretchr/testify/assert"
 )

  func TestCheckForWinner(t *testing.T) {
 	players = []string{"player1", "player2", "winner"}
 	allCards = make(map[string][]Card)
 	dealCards()

	assert.Equal(t, "", checkForWinner())
	 
 	allCards[players[2]] = make([]Card, 0)
 	assert.Equal(t, "winner", checkForWinner())
 } 