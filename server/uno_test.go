package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

//Test colors
func TestRandColor(t *testing.T) {
	assert.Equal(t,randColor(0),"red")
	assert.Equal(t,randColor(1),"blue")
	assert.Equal(t,randColor(2),"green")
	assert.Equal(t,randColor(3),"yellow")
}

//Test if card value is within the correct range
func TestNewRandomCard(t *testing.T){
	test := newRandomCard() 
	assert.Greater(t, test[0].Number, -1)
	assert.Less(t,test[0].Number,10)
}

//Test if card is in deck
func TestContains(t *testing.T){
	test := []string {"number", "reverse", "wild", "plus two"}
	index, isItThere := contains(test,"number")
	assert.Equal(t, index, 0)
	assert.Equal(t,isItThere,bool(true))
}

// Test if winner is found
func TestCheckWinner(t *testing.T){
	players = []string {"player1", "player2", "player3"}
	allCards = make(map[string][]Card)
	dealCards()
	assert.Equal(t, "", checkForWinner())
	allCards[players[0]] = make([]Card,0)
	assert.Equal(t, "player1", checkForWinner())
}

// Testing Check failure in TestCheckID
func TestCheckID(t *testing.T) {
	assert.Equal(t,checkID("user"),bool(false))
}