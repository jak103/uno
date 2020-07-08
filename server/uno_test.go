package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/labstack/echo"
)

func TestRandColor(t *testing.T) {
	// Setup
	assert.Equal(t,randColor(0),"red")
	assert.Equal(t,randColor(1),"blue")
	assert.Equal(t,randColor(2),"green")
	assert.Equal(t,randColor(3),"yellow")
}

func TestCheckID(t *testing.T) {
	fmt.Print("hello\n")
	assert.Equal(t,checkID("user"),bool(false))
}

func TestContains(t *testing.T){
	test := []string {"number", "reverse", "wild", "plus two"}
	index, isItThere := contains(test,"number")
	assert.Equal(t, index, 0)
	assert.Equal(t,isItThere,bool(true))
}

func TestNewRandomCard(t *testing.T){
	test := newRandomCard() 

	assert.Greater(t, test[0].Number, -1)
	assert.Less(t,test[0].Number,10)
}

// func TestCreateNewGame(t *testing.T){
// 	test := createNewGame(e)
// 	fmt.Print(test)
// }