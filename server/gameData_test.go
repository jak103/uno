package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCardCount(t *testing.T) {
	//TODO: Setup mock database

	//TODO: Add actual tests

	val, err := getCardCount("")
	assert.Equal(t, "", val)
	assert.Error(t, err)
}
