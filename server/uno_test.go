package main

import (
	"fmt"
	"reflect" // Used for checking types
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRandColor(t *testing.T) {
	// randColor should simply return string from int
	fmt.Println("Testing randColor()")
	assert.Equal(t, "red", randColor(0))
	assert.Equal(t, "blue", randColor(1))
	assert.Equal(t, "green", randColor(2))
	assert.Equal(t, "yellow", randColor(3))
	assert.Equal(t, "", randColor(-1))
	assert.Equal(t, "", randColor(4))
}

func TestNewRandomCard(t *testing.T) {
	fmt.Println("Testing newRandomCard()")
	for i := 0; i < 100; i++ { // Test 100 cards
		card := newRandomCard()
		inRange := card[0].Number < 10 && card[0].Number >= 0
		assert.Equal(t, true, inRange)
		// checking the types returned from creating a new card with reflect
		assert.Equal(t, reflect.TypeOf(0), reflect.TypeOf(card[0].Number))
		assert.Equal(t, reflect.TypeOf(""), reflect.TypeOf(card[0].Color))
	}

}

func TestContains(t *testing.T) {
    fmt.Println("Testing contains()")
    colors := []string{"red", "blue", "green", "yellow"}
    // If list contains string, should return (index, true)
    index, found := contains(colors, "yellow")
    assert.Equal(t, true, found)
    assert.Equal(t, 3, index)

    // If the list does not find string, should return (-1, false)
    index, found = contains(colors, "orange")
    assert.Equal(t, false, found)
    assert.Equal(t, -1, index)
}