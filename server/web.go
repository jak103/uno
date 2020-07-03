//This file will have all the data/functions to organize and return to the front-end
package main

import(
	"github.com/labstack/echo"
	"strconv"
	"fmt"
)

// Response is what the frontend will see
type Response struct {
	ValidGame bool                   `json:"valid"` // Valid game id
	Payload   map[string]interface{} `json:"payload"`
}

var db DB = DB{}

func getParams(c echo.Context) (bool, string, string) {
	return false, c.Param("game"), c.Param("username")
}

func withJSONPayload(gameCode string, username string) map[string]interface{} {
	payload := make(map[string]interface{})

	// Update known variables
	payload["game_id"] = gameCode
	payload["all_players"] = db.getAllPlayers(gameCode)
	payload["current_card"] = db.getCurrentCard(gameCode)
	payload["current_player"] = db.getCurrentPlayer(gameCode)
	tmp := db.getCurrentPlayerCards(gameCode, username)
	fmt.Println("Player Deck: ", tmp)
	payload["deck"] = db.getCurrentPlayerCards(gameCode, username)
	payload["game_over"] = checkForWinner(db.getAllCards(gameCode))

	fmt.Println("Payload: ", payload)
	return payload
}

func drawCard(c echo.Context) *Response {
	validGameID, id, user := getParams(c)

	if validGameID = db.isValidGame(id); validGameID && db.hasGameStarted(id) {
		deck := db.getDeckOfCards(id)
		card, _ := drawFromDeck(deck)
		db.removeCardFromDeck(card[0], id)
		return &Response{validGameID, withJSONPayload(id, user)}
	}
	return &Response{validGameID, nil}
}

func playCard(c echo.Context) *Response {
	num, _ := strconv.Atoi(c.Param("number"))
	color := c.Param("color")
	validGameID, id, user := getParams(c)
	fmt.Println("Valid Game: ", db.isValidGame(id))
	fmt.Println("hasGameStarted: ", db.hasGameStarted(id))
	fmt.Println("Number: ", num)
	fmt.Println("Color: ", color)
	if validGameID = db.isValidGame(id); validGameID && db.hasGameStarted(id) {
		currentCard := db.getCurrentCard(id)
		playerCard := Card{num, color}
		fmt.Println("New card to be played: ", playerCard)
		if canPlayNewCard(playerCard, currentCard[0]) {
			fmt.Println("Can play the card")
			db.setCurrentCard([]Card{playerCard})
			db.removeCardFromPlayersDeck(playerCard, id, user)
			
			return &Response{validGameID, withJSONPayload(id, user)}
		} else {
			fmt.Println("Could not play card")
		}
	}
	return &Response{validGameID, nil}
}

func dealCardsAndStartGame(c echo.Context) *Response {
	validGameID, id, user := getParams(c)

	if validGameID = db.isValidGame(id); validGameID {
		allCards := createDeck()
		fmt.Println("All Cards: ", allCards)
		playerCards := db.getAllCards(id)
		fmt.Println("Player cards: ", playerCards)
		topCard := dealCards(&playerCards, &allCards)
		fmt.Println("All Cards: ", allCards)
		fmt.Println("Player cards: ", playerCards)

		db.updatePlayerCards(playerCards, id)
		db.setDeckOfCardsForDrawPile(allCards, id)
		db.setCurrentCard(topCard)
		db.startGame(id)
		return &Response{validGameID, withJSONPayload(id, user)}
	}
	return &Response{validGameID, nil}	
}

func heartBeat(c echo.Context) *Response {
	validGameID, id, user := getParams(c)

	if validGameID = db.isValidGame(id); validGameID && db.hasGameStarted(id) {
		return &Response{validGameID, withJSONPayload(id, user)}
	}
	return &Response{validGameID, nil}
}

func joinGameIfValid(c echo.Context) *Response { // If user can join
	validGameID, id, user := getParams(c)

	if validGameID = db.isValidGame(id); validGameID {
		db.addUserToGame(user)
		return &Response{validGameID, withJSONPayload(id, user)}
	}
	return &Response{validGameID, nil}
}

func createNewGame(c echo.Context) *Response {
	id := db.addNewGame()
	user := c.Param("username")
	return &Response{true, withJSONPayload(id, user)}
}