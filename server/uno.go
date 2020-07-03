package main

import (
	"math/rand"
	"time"

	// "cloud.google.com/go/firestore"
	// "google.golang.org/api/iterator"
)

// Card represents and uno card, and can be represented as json for frontend
type Card struct {
	Number int    `json:"number"`
	Color  string `json:"color"`
}

func contains(arr []string, val string) (int, bool) {
	for i, item := range arr {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func randColor(i int) string {
	switch i {
	case 0:
		return "red"
	case 1:
		return "blue"
	case 2:
		return "green"
	case 3:
		return "yellow"
	}
	return ""
}

func canPlayNewCard(newCard, currentCard Card) bool {
	if newCard.Color == currentCard.Color || newCard.Number == currentCard.Number {
		return true
	}
	return false
}

func drawFromDeck(deck []Card) ([]Card, int) {
	newCardIndex := rand.Intn(len(deck))
	cardDrawn := []Card{deck[newCardIndex]}

	return cardDrawn, newCardIndex
}

// TODO: need to deal the actual cards, not just random numbers
func dealCards(playersCards *map[string][]Card, drawPileCards *[]Card) []Card {
	for player := range *(playersCards) {
		for i := 0; i < 7; i++ {
			// TODO: Test this and make sure its working
			index := rand.Intn(len(*(drawPileCards)))
			randomCard := (*drawPileCards)[index]
			(*playersCards)[player] = append((*playersCards)[player], randomCard)
			
			(*drawPileCards)[len((*drawPileCards))-1], (*drawPileCards)[index] = (*drawPileCards)[index], (*drawPileCards)[len((*drawPileCards))-1]
			(*drawPileCards) = (*drawPileCards)[:len((*drawPileCards))-1]
		}
	}

	return []Card{(*drawPileCards)[0]} // TODO: Create function that will remove a card from a deck
}


func checkForWinner(gameCards map[string][]Card) string {
	for player := range gameCards {
		if len(gameCards[player]) == 0 {
			return player
		}	
	}
	return ""
}

func createDeck() []Card {
	colors := []string{"yellow", "green", "blue", "red"}
	deck := []Card{}
	for _, color := range(colors) {
		for i := 1; i < 10; i++ {
			deck = append(deck, Card{i, color});
		} 
	}

	rand.Seed(time.Now().UnixNano())
	for i := len(deck) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}

	return deck
}
