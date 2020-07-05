package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandColor(t *testing.T) {
	testRed := randColor(0)    // <0> should return "red"
	testBlue := randColor(1)   // <1> should return "blue"
	testGreen := randColor(2)  // <2> should return "green"
	testYellow := randColor(3) // <3> should return "yellow"

	assert.Equal(t, testRed, "red")
	assert.Equal(t, testBlue, "blue")
	assert.Equal(t, testGreen, "green")
	assert.Equal(t, testYellow, "yellow")
}
