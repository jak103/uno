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

	// expected, actual, msg
	assert.NotEqual(r2, r1, "The two colors should be the same.")
}

func TestDefaultPayload(t *testing.T) {
	fmt.Println("Running Payload Test")
	assert := assert.New(t)

	nilPayload := newPayload("user_test")

	var emptyCardArray []Card
	var emptyPlayerArray []string = []string{}

	assert.Equal(emptyCardArray, nilPayload["current_card"], "Current card should be empty")
	assert.Equal("", nilPayload["current_player"], "Player should be empty")
	assert.Equal(emptyPlayerArray, nilPayload["all_players"], "Players list should be empty")
	assert.Equal(emptyCardArray, nilPayload["deck"], "Card should be empty")
	assert.Equal("", nilPayload["game_id"], "Game id should be empty")
	assert.Equal("", nilPayload["game_over"], "Game over should be empty")
}
