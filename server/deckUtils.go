package main

import (
	"fmt"
	"time"
	"math/rand"
)

// Represents a card
// Uses Value instead of Number
// to more accurately represent non-numerical cards
type UnoCard struct {
	Color string `json:"color"`
	Value string `json:"value"`
}

// Prints a card's color and value
// Useful during development
func printCard(card UnoCard) {
	fmt.Println(card.Color, card.Value)
}

// Prints a slice of cards
// Useful during development
func printCards(cards []UnoCard) {
	for _, card := range cards {
		printCard(card)
	}
}

// Returns the standard card colors, standard card counts, and wild card counts
// It would be awesome to have this be customizable from the front-end
// so you can play with 15 reverse cards if you want
// or with 6 colors or something
func getDeckConfig() ([4]string, map[string]int, map[string]int) {
	colors := [4]string{"red", "blue", "green", "yellow"}

	standardCardCounts := map[string]int {
		"zero": 1,
		"one": 2,
		"two": 2,
		"three": 2,
		"four": 2,
		"five": 2,
		"six": 2,
		"seven": 2,
		"eight": 2,
		"nine": 2,
		"skip": 2,
		"draw_two": 2,
		"reverse": 2,
	}

	wildCardCounts := map[string]int {
		"wild": 4,
		"wild_draw_four": 4,
	}

	return colors, standardCardCounts, wildCardCounts
}

// Returns the cards provided, but in a random order
// Credit to https://yourbasic.org/golang/shuffle-slice-array/
func shuffleCards(a []UnoCard) []UnoCard {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}

// Generates a deck with standard card values and counts for each color
// and with wild card values and counts as well.
// Shuffles the deck before returning it.
// This function is not necessarily efficient - feel free to optimize.
func generateShuffledDeck() []UnoCard {
	colors, standardCardCounts, wildCardCounts := getDeckConfig()
	deck := []UnoCard{}
	for cardValue, count := range standardCardCounts {
		for i := 0; i < count; i++ {
			for _, color := range colors {
				deck = append(deck, UnoCard{color, cardValue})
			}
		}
	}

	for cardValue, count := range wildCardCounts {
		for i := 0; i < count; i++ {
			deck = append(deck, UnoCard{"wild", cardValue})
		}
	}

	return shuffleCards(deck)
}
