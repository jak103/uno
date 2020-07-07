package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	assert.Equal(t, 12, add(8, 4))
}

func TestSubtract(t *testing.T) {
	assert.Equal(t, 6, subtract(9, 3))
}

func TestMultiple(t *testing.T) {
	assert.Equal(t, 100, multiply(10, 10))
}

func TestDivide(t *testing.T) {
	assert.Equal(t, 10, divide(100, 10))
}

func TestMain(t *testing.T) {
	main()
}
