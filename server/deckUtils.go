package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jak103/uno/model"
)

// Prints a card's color and value
// Useful during development
func printCard(card model.Card) {
	fmt.Println(card.Color, card.Value)
}

// Prints a slice of cards
// Useful during development
func printCards(cards []model.Card) {
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

	standardCardCounts := map[string]int{
		"0":  1,
		"1":  2,
		"2":  2,
		"3":  2,
		"4":  2,
		"5":  2,
		"6":  2,
		"7":  2,
		"8":  2,
		"9":  2,
		"S":  2,
		"D2": 2,
		"R":  2,
	}

	wildCardCounts := map[string]int{
		"W":  4,
		"W4": 4,
	}

	return colors, standardCardCounts, wildCardCounts
}

// Returns the cards provided, but in a random order
// Credit to https://yourbasic.org/golang/shuffle-slice-array/
func shuffleCards(a []model.Card) []model.Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}

// Generates a deck with standard card values and counts for each color
// and with wild card values and counts as well.
// Shuffles the deck before returning it.
// This function is not necessarily efficient - feel free to optimize.
func generateShuffledDeck() []model.Card {
	colors, standardCardCounts, wildCardCounts := getDeckConfig()
	deck := []model.Card{}
	for cardValue, count := range standardCardCounts {
		for i := 0; i < count; i++ {
			for _, color := range colors {
				deck = append(deck, model.Card{color, cardValue})
			}
		}
	}

	for cardValue, count := range wildCardCounts {
		for i := 0; i < count; i++ {
			deck = append(deck, model.Card{"wild", cardValue})
		}
	}

	return shuffleCards(deck)
}
