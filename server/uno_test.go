package main

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

// Utility Functions For Testing
func ValidColors(cardColor string) bool {
  if (cardColor == "red" || cardColor == "yellow" || cardColor == "green" || cardColor == "blue") {
	return true
  } else{
	return false
  }
}

func ValidNum(num int) bool {
  var retVal bool = false
  switch num {
	case 0: retVal = true
	case 1: retVal = true
	case 2: retVal = true
	case 3: retVal = true
	case 4: retVal = true
	case 5: retVal = true
	case 6: retVal = true
	case 7: retVal = true
	case 8: retVal = true
	case 9: retVal = true
  }
  return retVal	
}

// Unit Tests
func TestRandColor(t *testing.T) {
  for i := 100; i > 0; i-- {
 	testVal := newRandomCard()
	assert.True(t, ValidNum(testVal[0].Number))
	assert.True(t, ValidColors(testVal[0].Color))
	assert.Equal(t, 1, len(testVal))
  }
}

func TestNewPayload(t *testing.T) {
  testVal := newPayload("")
  assert.Equal(t, []string{} ,testVal["all_players"])
  assert.Equal(t, "12234" ,testVal["game_id"])
}

func TestId(t *testing.T) {
	assert.True(t, checkID("12234"))
}

func TestContains(t *testing.T) {
	var players []string = []string{}
	players = append(players, "one")

	a, b := contains(players, "one")
	assert.Equal(t, 0, a)
	assert.True(t, b)
}
