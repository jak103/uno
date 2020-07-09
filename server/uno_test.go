package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestRandColor(t *testing.T) {
	assert.Equal(t, "red", randColor(0))
	assert.Equal(t, "blue", randColor(1))
	assert.Equal(t, "green", randColor(2))
	assert.Equal(t, "yellow", randColor(3))

}

func TestCheckID(t *testing.T) {
	assert.Equal(t, true, checkID("12234"))
}

func testGameStarted(t *testing.T) {
	assert.Equal(t, false, gameStarted)
}

func testContains(t *testing.T) {
	var testArray []string = []string{"a", "b", "c"}

	intResult, boolResult := contains(testArray, "hi")

	assert.Equal(t, -1, intResult)
	assert.Equal(t, false, boolResult)

	intResult, boolResult = contains(testArray, "a")

	assert.Equal(t, 0, intResult)
	assert.Equal(t, true, boolResult)

	intResult, boolResult = contains(testArray, "b")

	assert.Equal(t, 1, intResult)
	assert.Equal(t, true, boolResult)

	intResult, boolResult = contains(testArray, "c")

	assert.Equal(t, 2, intResult)
	assert.Equal(t, true, boolResult)
}

func testCheckForWinner(t *testing.T) {
	assert.Equal(t, "", checkForWinner())
}
