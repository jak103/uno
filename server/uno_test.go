package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCheckID(t *testing.T) {
	id := "tester"
	assert.Equal(t, false, checkID(id))
}

func TestRandomColor(t *testing.T) {
	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			assert.Equal(t, "red", randColor(i))
		case 1:
			assert.Equal(t, "blue", randColor(i))		
		case 2:
			assert.Equal(t, "green", randColor(i))
		case 3:
			assert.Equal(t, "yellow", randColor(i))
		}
	}
}

func TestContainsString(t *testing.T) {
	id := "Tester"
	list := []string {"Player", "Loser", "Tester", "Winner"}
	outputIndex, outputBool := contains(list, id)
	assert.Equal(t, outputIndex, 2)
	assert.Equal(t, outputBool, true)
}