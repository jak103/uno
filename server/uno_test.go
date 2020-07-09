package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Testing Random Colors
func TestRandColor(t *testing.T) {
	assert.Equal(t,randColor(0),"red")
	assert.Equal(t,randColor(1),"blue")
	assert.Equal(t,randColor(2),"green")
	assert.Equal(t,randColor(3),"yellow")
}

// Testing Check failure in TestCheckID
func TestCheckID(t *testing.T) {
	assert.Equal(t,checkID("user"),bool(false))
}

//Testing Contains Method. Seeing if a given card is found
func TestContains(t *testing.T){
	test := []string {"number", "reverse", "wild", "plus two"}
	index, isItThere := contains(test,"number")
	assert.Equal(t, index, 0)
	assert.Equal(t,isItThere,bool(true))
}

//Testing if randcom card generated is less than 10 and greater than -1
func TestNewRandomCard(t *testing.T){
	test := newRandomCard() 

	assert.Greater(t, test[0].Number, -1)
	assert.Less(t,test[0].Number,10)
}

// Testing if Winners are found.
func TestCheckWinner(t *testing.T){
	players = []string {"player1", "player2", "player3"}
	allCards = make(map[string][]Card)

	dealCards()

	assert.Equal(t, "", checkForWinner())

	allCards[players[0]] = make([]Card,0)

	assert.Equal(t, "player1", checkForWinner())
}