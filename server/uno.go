package main

import (
	"github.com/labstack/echo"
	"math/rand"
)


////////////////////////////////////////////////////////////
// Structs used for the talking with frontend
////////////////////////////////////////////////////////////
type Response struct {
	ValidGame bool 				   `json:"valid"`  // Valid game id
	Payload map[string]interface{} `json:"payload"`
}

type Card struct {
	Number int 	 `json:"number"`
	Color string `json:"color"`
}


////////////////////////////////////////////////////////////
// Utility functions used in place of firebase
////////////////////////////////////////////////////////////
func randColor(i int) string {
	switch (i) {
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


////////////////////////////////////////////////////////////
// All the data needed for a simulation of the game
// eventually, this will be replaced with firebase
////////////////////////////////////////////////////////////
var gameID string = ""
var currCard []Card = nil   // The cards are much easier to render as a list
var players []string = []string{"Bill", "Bob", "Joe"}
var playerIndex = rand.Intn(len(players))  // Used to iterate through the players
var currPlayer string = players[playerIndex]
var allCards map[string][]Card = make(map[string][]Card) // k: username, v: list of cards
var gameStarted bool = false


////////////////////////////////////////////////////////////
// Utility functions
////////////////////////////////////////////////////////////
func newRandomCard() []Card {
	return []Card{Card{rand.Intn(9), randColor(rand.Intn(3))}}
}

func newPayload(user string) map[string]interface{} { // User will default to "" if not passed
	payload := make(map[string]interface{})
	

	// Update known variables 
	payload["current_card"] = currCard
	payload["current_player"] = currPlayer
	payload["all_players"] = players
	payload["deck"] = allCards[user] // returns nil if currPlayer = "" or user not in allCards
	payload["game_id"] = gameID
	
	return payload
}

func checkID(id string) bool {
	return id == gameID
}

func contains(arr []string, val string) (int, bool) {
	for i, item := range arr {
		if item == val {
			return i, true
		}
	}
	return -1, false
}



////////////////////////////////////////////////////////////
// These are all of the functions for the game -> essentially public functions
////////////////////////////////////////////////////////////
func updateGame(c echo.Context) *Response {
	success := false
	if success = checkID(c.Param("game")); success && gameStarted {
		return &Response{true, newPayload(c.Param("username"))}
	}
	return &Response{false, nil}
}

func createNewGame(c echo.Context) *Response {
	gameID = "12234"
	return &Response{true, newPayload("")}
}

func joinGame(c echo.Context) *Response {
	if checkID(c.Param("game")) {
		user := c.Param("username")
		if _, found := contains(players, user); !found {
			players = append(players, user)
			allCards[user] = nil // No cards yet
		}
		return &Response{true, newPayload(c.Param("username"))}
	}
	return &Response{false, nil}  // bad game_id
}

func playCard(c Card, r *Response) bool {
	success := false
	if c.Color == currCard[0].Color  || c.Number == currCard[0].Number {
		playerIndex = (playerIndex + 1) % len(players)
		currPlayer = players[playerIndex]
		currCard[0] = c

		r.Payload["current_card"] = currCard
		r.Payload["current_player"] = currPlayer
		success = true
	}

	return success
}

// TODO: Keep track of current card that is top of the deck
func drawCard(r *Response) {
	r.Payload["current_card"] = newRandomCard()
}

// TODO: need to deal the actual cards, not just random numbers
func dealCards() {
	// The game has started, no more players are joining
	// loop through players, set their cards
	gameStarted = true

	for k := range players {
		cards := []Card{}
		for i := 0; i < 7; i++ {
			cards = append(cards, Card{rand.Intn(9), randColor(rand.Intn(3))})
		}
		allCards[players[k]] = cards
	}

	currCard = []Card{Card{rand.Intn(9), randColor(rand.Intn(3))}}
}
