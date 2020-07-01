//This file will have all the data/functions to organize and return to the front-end
package main

import(
	"github.com/labstack/echo"
	"strconv"
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
	payload["deck"] = db.getCurrentPlayerCards(gameCode, username)
	payload["game_over"] = checkForWinner(db.getAllCards(gameCode))

	return payload
}

func drawCard(c echo.Context) *Response {
	validGameID, id, user := getParams(c)

	if validGameID = db.isValidGame(id); validGameID && db.hasGameStarted(id) {
		
		return &Response{validGameID, withJSONPayload(id, user)}
	}
	return &Response{validGameID, nil}
}

func playCard(c echo.Context) *Response {
	num, _ := strconv.Atoi(c.Param("number"))
	color := c.Param("Color")
	validGameID, id, user := getParams(c)

	if validGameID = db.isValidGame(id); validGameID && db.hasGameStarted(id) {
		currentCard := db.getCurrentCard(id)
		playerCard := Card{num, color}

		if canPlayNewCard(playerCard, currentCard[0]) {
			db.setCurrentCard([]Card{playerCard})
			db.removeCardFromPlayersDeck(playerCard, id, user)
			
			return &Response{validGameID, withJSONPayload(id, user)}
		}
	}
	return &Response{validGameID, nil}
}

func dealCardsAndStartGame(c echo.Context) *Response {
	validGameID, id, user := getParams(c)

	if validGameID = db.isValidGame(id); validGameID {
		dealCards(id)
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