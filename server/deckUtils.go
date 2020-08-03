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

// Get's the number of decks to use based on the number of players
func numDecksToUse(numPlayers int) int {
	return numPlayers/5 + 1
}

// Returns the card colors, card counts, and wild card counts based
// on the number of players in the game
func getDeckConfigByPlayerSize(numDecks int) ([4]string, map[string]int, map[string]int) {
	colors := [4]string{"red", "blue", "green", "yellow"}

	standardCardCounts := map[string]int{
		"0":  1 * numDecks,
		"1":  2 * numDecks,
		"2":  2 * numDecks,
		"3":  2 * numDecks,
		"4":  2 * numDecks,
		"5":  2 * numDecks,
		"6":  2 * numDecks,
		"7":  2 * numDecks,
		"8":  2 * numDecks,
		"9":  2 * numDecks,
		"S":  2 * numDecks,
		"D2": 2 * numDecks,
		"R":  2 * numDecks,
	}

	wildCardCounts := map[string]int{
		"W":  4 * numDecks,
		"W4": 4 * numDecks,
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
func generateShuffledDeck(numPlayers int) []model.Card {
	numDecks := numDecksToUse(numPlayers)
	colors, standardCardCounts, wildCardCounts := getDeckConfigByPlayerSize(numDecks)
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
			deck = append(deck, model.Card{"black", cardValue})
		}
	}

	return shuffleCards(deck)
}
