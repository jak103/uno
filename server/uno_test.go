package main

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRandColor(t *testing.T) {
	fmt.Println("Testing Rand Color")
	testRed := randColor(0)
	assert.Equal(t, testRed, "red")
	testBlue := randColor(1)
	assert.Equal(t, testBlue, "blue")
	testGreen := randColor(2)
	assert.Equal(t, testGreen, "green")
	testYellow := randColor(3)
	assert.Equal(t, testYellow, "yellow")
	testOverflow := randColor(4)
	assert.Equal(t, testOverflow, "")
}

func TestContains(t *testing.T) {
	array := []string{"a", "b", "c", "d", "e"}
	index, doesContain := contains(array, "a")
	assert.Equal(t, index, 0)
	assert.True(t, doesContain)
}

func TestNewRandomCard(t *testing.T) {
	card := newRandomCard()
	assert.Less(t, card[0].Number, 10)
	assert.Greater(t, card[0].Number, -1)
}

