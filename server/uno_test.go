package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const ITERATIONS = 10

func isValidColor(input string) bool {
	switch input {
	case
		"red",
		"blue",
		"green",
		"yellow":
		return true
	}
	return false
}

func isValidNumber(input int) bool {
	switch input {
	case
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
		return true
	}
	return false
}

func TestNewRandomCard(t *testing.T) {
	for i := 0; i < ITERATIONS; i++ {
		card := newRandomCard()
		assert.Equal(t, 1, len(card))
		assert.True(t, isValidColor(card[0].Color))
		assert.True(t, isValidNumber(card[0].Number))
	}
}

func TestRandColor(t *testing.T) {
	for i := 0; i < ITERATIONS; i++ {
		assert.True(t, isValidColor(randColor(i%4)))
	}
}

func TestCheckID(t *testing.T) {
	createNewGame()
	assert.Equal(t, true, checkID("12234"))
}

func TestContains(t *testing.T) {
	var EMPTY = []string{}
	var ALPHA = []string{"one value"}
	var BETA = []string{"first", "second"}

	var index int
	var isFound bool
	index, isFound = contains(EMPTY, "")
	assert.False(t, isFound)
	assert.Equal(t, -1, index)

	index, isFound = contains(EMPTY, "something")
	assert.False(t, isFound)
	assert.Equal(t, -1, index)

	index, isFound = contains(ALPHA, "")
	assert.False(t, isFound)
	assert.Equal(t, -1, index)

	index, isFound = contains(ALPHA, "something")
	assert.False(t, isFound)
	assert.Equal(t, -1, index)

	index, isFound = contains(ALPHA, "one value")
	assert.True(t, isFound)
	assert.Equal(t, 0, index)

	index, isFound = contains(BETA, "")
	assert.False(t, isFound)
	assert.Equal(t, -1, index)

	index, isFound = contains(BETA, "something")
	assert.False(t, isFound)
	assert.Equal(t, -1, index)

	index, isFound = contains(BETA, "second")
	assert.True(t, isFound)
	assert.Equal(t, 1, index)
}
