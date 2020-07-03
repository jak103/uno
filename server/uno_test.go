package main

import (
	"testing"
)

func TestRandColor(t *testing.T) {

	red := randColor(0)
	blue := randColor(1)
	green := randColor(2)
	yellow := randColor(3)
	other1 := randColor(-2)
	other2 := randColor(25)

	if red != "red" {
		t.Error(red + " != red")
	}
	if blue != "blue" {
		t.Error(blue + " != blue")
	}
	if green != "green" {
		t.Error(green + " != green")
	}
	if yellow != "yellow" {
		t.Error(yellow + " != yellow")
	}
	if other1 != "" {
		t.Error(other1 + " != ")
	}
	if other2 != "" {
		t.Error(other2 + " != ")
	}
}
