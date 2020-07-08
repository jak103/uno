package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


/*Example formatting of a test from class sample code
func TestSubtract(t *testing.T) {
	assert.Equal(t, 6, subtract(9, 3))
}
*/

/******************************************************************************
This function will test the function randColor in uno.go
		randColor can return "red", "blue", "green", "yellow" or ""
******************************************************************************/
func TestRandColor(t *testing.T) {
	
	assert.Equal(t, "red", randColor(0))
	assert.Equal(t, "blue", randColor(1))
	assert.Equal(t, "green", randColor(2))
	assert.Equal(t, "yellow", randColor(3))
	assert.Equal(t, "", randColor(4))

}


/******************************************************************************
This function will test the function newRandomCard in uno.go
		newRandomCard can return a random colored card numbered between 0 to 9
******************************************************************************/
func TestNewRandomCard(t *testing.T){

	testCard := newRandomCard()[0] 

	assert.Greater(t, testCard.Number, -1)
	assert.Less(t, testCard.Number, 10)
	

}

/******************************************************************************
This function will test the function checkID in uno.go
		right now it is hard coded to give the id 12234
******************************************************************************/
func TestCheckID(t *testing.T){

	assert.Equal(t, true, checkID("12234"))

}


/******************************************************************************
This function will test the function contains in uno.go
		contains is given an array and a value, and will check the array for
		that value - returning the index and true if it found it.
		if the value is not found in the array it returns -1 and false
		--- currently our cards are random, but this could be useful later for
			testing if our deck contains all the cards an uno deck should have.
******************************************************************************/
func TestContains(t *testing.T){

	deck := []string{"Red1", "Red2", "Red3", "Red4", "Red5", "Red6", "Red7", "Red8", "Red9"}
	testCardIndex, testCardBool := contains(deck, "Red1")

	assert.True(t, testCardBool)
	assert.Equal(t, testCardIndex, 0)

}
