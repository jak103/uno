package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandColor(t *testing.T) {
	/*

	   the 'randColor' function needs to return a string
	   value that is either "red", "blue", "green", or "yellow" for valid
	   input and return an empty string ("") for invalid input

	*/

	testRed := randColor(0)    // <0> should return "red"
	testBlue := randColor(1)   // <1> should return "blue"
	testGreen := randColor(2)  // <2> should return "green"
	testYellow := randColor(3) // <3> should return "yellow"
	testBad := randColor(4)    // <4> should return ""

	assert.Equal(t, testRed, "red")
	assert.Equal(t, testBlue, "blue")
	assert.Equal(t, testGreen, "green")
	assert.Equal(t, testYellow, "yellow")
	assert.Equal(t, testBad, "")
}

func TestNewRandomCard(t *testing.T) {
	/*
	   the 'newRandomCard' func needs to return a card with
	   an integer value no less than 0 and no greater than 9.

	   the 'newRandomCard' function also needs to return a string
	   value that is either "red", "blue", "green", or "yellow" for valid
	   input and return an empty string ("") for invalid input, which is
	   tested in 'TestRandColor'
	*/

	myCard := newRandomCard()[0]

	assert.Less(t, myCard.Number, 10)
	assert.Greater(t, myCard.Number, -1)

}

func TestContains(t *testing.T) {

	assert.Equal(t, "0", "0")

	hasRed := []string{"blue", "green", "red", "yellow"}
	//noRed := []string{"blue", "green", "orange", "yellow"}
	//notExactlyRed := []string{" red", "red ", "rred", "RED"}

	hasRed_intVal, hasRed_boolVal := contains(hasRed, "red")
	//noRed_intVal, noRed_boolVal := contains(noRed, "red")
	//notExactlyRed_intVal, notExactlyRed_boolVal := contains(notExactlyRed, "red")

	assert.Equal(t, hasRed_intVal, 2) // "red" is in index 2 of array
	assert.Equal(t, hasRed_boolVal, true)

}
