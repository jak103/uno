package main

import (
	"github.com/labstack/echo"
	"math/rand"
	"fmt"
)


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
var gameID string = "12234"
var currCard []*Card = []*Card{&Card{5, "red"}}  // The cards are much easier to render as a list
var players []string = []string{"Bill", "Bob", "Jill"}
var playerIndex = rand.Intn(len(players))
var currPlayer string = players[playerIndex]


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
// Utility functions
////////////////////////////////////////////////////////////
func newRandomCard() []*Card {
	return []*Card{&Card{rand.Intn(9), randColor(rand.Intn(3))}}
}

func newPayload() map[string]interface{} {
	return make(map[string]interface{})
}

func checkID(id string) bool {
	return id == gameID
}

func createNewGame(c echo.Context) *Response {
	payload := newPayload()
	payload["game_id"] = gameID
	return &Response{true, payload}
}

func contains(arr []string, val string) (int, bool) {
	for i, item := range arr {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func joinGame(c echo.Context) *Response {
	if checkID(c.Param("game")) {
		if _, found :=contains(players, c.Param("username")); !found {
			players = append(players, c.Param("username"))
		}
		fmt.Println(players)
		return &Response{true, nil}	
	}
	return &Response{false, nil}
}

func newResponse(c echo.Context) *Response {
	payload := make(map[string]interface{})
	payload["current_card"] = currCard
	payload["current_player"] = currPlayer
	payload["players"] = players
	return &Response{checkID(c.Param("game")), payload}
}


////////////////////////////////////////////////////////////
// These are all of the functions for the game
////////////////////////////////////////////////////////////
func playCard(c *Card, r *Response) bool {
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

func drawCard(r *Response) {
	r.Payload["current_card"] = newRandomCard()
}

func dealCards(r *Response) {
	cards := []*Card{}
	for i := 0; i < 7; i++ {
		cards = append(cards, &Card{rand.Intn(9), randColor(rand.Intn(3))})
	}
	
	r.Payload["deck"] = cards
}
