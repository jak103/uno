package main

import (
	"testing"
)

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

func TestDrawFromDeck(t *testing.T) {
	deck := []Card{
		Card{5, "red"},
		Card{3, "green"},
		Card{2, "yellow"},
		Card{6, "blue"},
		Card{9, "red"},
	}

	if length := len(deck); length != 5 {
		t.Error("Deck should be length 5, but failed, got: ", length)
	}

	_, index := drawFromDeck(deck)
	deck[len(deck)-1], deck[index] = deck[index], deck[len(deck)-1]
	deck = deck[:len(deck)-1]
	if length := len(deck); length != 4 {
		t.Error("Deck should be length 4, but failed, got: ", length)
	}

	_, index = drawFromDeck(deck)
	deck[len(deck)-1], deck[index] = deck[index], deck[len(deck)-1]
	deck = deck[:len(deck)-1]
	if length := len(deck); length != 3 {
		t.Error("Deck should be length 3, but failed, got: ", length)
	}

	_, index = drawFromDeck(deck)
	deck[len(deck)-1], deck[index] = deck[index], deck[len(deck)-1]
	deck = deck[:len(deck)-1]
	if length := len(deck); length != 2 {
		t.Error("Deck should be length 2, but failed, got: ", length)
	}
}