package main

import (
	"testing"
)

func TestColors(t *testing.T) {
	red := 0
	blue := 1
	green := 2
	yellow := 3

	test := randColor(red)
	if test != "red" {
        t.Error("Color test: red..got: ", test)
	}

	test = randColor(blue)
	if test != "blue" {
        t.Error("Color test: blue..got: ", test)
	}

	test = randColor(green)
	if test != "green" {
        t.Error("Color test: green..got: ", test)
	}

	test = randColor(yellow)
	if test != "yellow" {
        t.Error("Color test: yellow..got: ", test)
	}
}

func TestContains(t *testing.T) {
	var players []string = []string{"one", "two", "three", "four"};

	if _, has := contains(players, "one"); !has {
		t.Error("Array should contain: \"one\", but returned false");
	}

	if _, has := contains(players, "two"); !has {
		t.Error("Array should contain: \"two\", but returned false");
	}

	if _, has := contains(players, "four"); !has {
		t.Error("Array should contain: \"four\", but returned false");
	}

	if _, has := contains(players, "seven"); has {
		t.Error("Array should not contain: \"seven\", but returned true");
	}
}