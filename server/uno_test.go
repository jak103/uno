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
	/*
		The 'contains' func takes in both an array and a value as a paramenter.
		If it finds the value in the array, it will return both the index where
		the value is as well as the boolean value 'true'. If the value is not in
		the array it will return both the integer value '-1' instead of an index
		as well as the boolean value 'false'
	*/

	// this tests the 'contains' funct with an array that contains the value
	hasRed := []string{"blue", "green", "red", "yellow"}
	hasRedIntVal, hasRedBoolVal := contains(hasRed, "red")
	assert.Equal(t, hasRedIntVal, 2) // "red" is in index 2 of array
	assert.Equal(t, hasRedBoolVal, true)

	// this tests the 'contains' funct with an array that doesn't contain the value
	noRed := []string{"blue", "green", "orange", "yellow"}
	noRedIntVal, noRedBoolVal := contains(noRed, "red")
	assert.Equal(t, noRedIntVal, -1) // returns -1 if value not found
	assert.Equal(t, noRedBoolVal, false)

	// this tests the 'contains' funct with an array that 'almost' contains the value
	notExactlyRed := []string{" red", "red ", "rred", "RED"}
	notExactlyRedIntVal, notExactlyRedBoolVal := contains(notExactlyRed, "red")
	assert.Equal(t, notExactlyRedIntVal, -1) // returns -1 if value not found
	assert.Equal(t, notExactlyRedBoolVal, false)

}
