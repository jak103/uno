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

