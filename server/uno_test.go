package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardDraw(t *testing.T) {
	card := newRandomCard() //assigns new card
	assert.Less(t, card[0].Number, 10)
	assert.Greater(t, card[0].Number, -1) //checks if the cards are within correct range

	assert.NotEmpty(t, card[0].Color) //makes sure color is getting assigned

	if card[0].Color == "green" { //checks for all the possible correct colors
		assert.Equal(t, card[0].Color, "green")

	} else if card[0].Color == "red" {
		assert.Equal(t, card[0].Color, "red")

	} else if card[0].Color == "blue" {
		assert.Equal(t, card[0].Color, "blue")

	} else if card[0].Color == "yellow" {
		assert.Equal(t, card[0].Color, "yellow")
	} else {
		assert.Fail(t, "color not correct") //if the "color" is not one of the good ones it will fail
	}

}
