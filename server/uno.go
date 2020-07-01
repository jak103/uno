package main

import (
	"math/rand"

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

func canPlayNewCard(newCard Card, currentCard Card) bool {
	if newCard.Color == currentCard.Color || newCard.Number == currentCard.Number {
		return true
	}
	return false
}

func drawFromDeck(deck []Card) []Card {
	newCardIndex := rand.Intn(len(deck))
	cardDrawn := []Card{deck[newCardIndex]}

	//Remove from deck: https://yourbasic.org/golang/delete-element-slice/
	go func() {
		deck[newCardIndex] = deck[len(deck)-1]
		deck[len(deck)-1] = Card{}
		deck = deck[:len(deck)-1]
	}() // This starts the thread
	
	return cardDrawn
}

// TODO: need to deal the actual cards, not just random numbers
func dealCards(gameCode string) {
	cards := db.getAllCards(gameCode)

	for player, deck := range cards {
		for i := 0; i < 7; i++ {
			deck = append(deck, Card{rand.Intn(10), randColor(rand.Intn(4))})
		}
	}

	db.updateCards(cards);
}


func checkForWinner(gameCards map[string][]Card) string {
	for player, cards := range gameCards {
		if len(gameCards[player]) == 0 {
			return player
		}	
	}
	return ""
}
