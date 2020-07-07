package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -cover

func TestColor(t *testing.T) {
	fmt.Println("Running Rand Color Test")
	assert := assert.New(t)

	r1 := randColor(0)
	r2 := ""

	assert.NotEqual(r1, r2, "The two colors should be the same.")
}

func TestPayload(t *testing.T) {
	fmt.Println("Running Payload Test")

	// nilPayload := newPayload("user_test")

	// payload["current_card"] = currCard
	// payload["current_player"] = currPlayer
	// payload["all_players"] = players
	// payload["deck"] = allCards[user] // returns nil if currPlayer = "" or user not in allCards
	// payload["game_id"] = gameID
	// payload["game_over"] = checkForWinner()
	// var emptyCardArray []Card
	// assert.Equal(nilPayload["current_card"], emptyCardArray, "Card should be nil")
	// assert.Equal(nilPayload["current_player"], nil, "Player should be nil")
	// assert.Equal(nilPayload["all_players"], nil, "All players should be nil")
	// assert.Equal(nilPayload["deck"], nil, "Card should be nil")
	// assert.Equal(nilPayload["game_id"], nil, "Card should be nil")
	// assert.Equal(nilPayload["game_over"], nil, "Card should be nil")
}
