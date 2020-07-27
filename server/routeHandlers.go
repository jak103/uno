package main

import (
	"net/http"

	"github.com/jak103/uno/model"
	"github.com/labstack/echo/v4"
)

var sim bool = true

// Response is a JSON structure which tells the front end about the game state
type Response struct {
	ValidGame bool                   `json:"valid"` // Valid game id
	Payload   map[string]interface{} `json:"payload"`
}

// Status is a JSON structure for sending general messages to the front end
type Status struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func setupRoutes(e *echo.Echo) {
	e.GET("/newgame", newGame)
	e.GET("/update/:game/:username", update)
	e.GET("/cardcount/:game", cardCount)
	e.POST("/startgame/:game/:username", startGame)
	e.POST("/login/:game/:username", login)
	e.POST("/play/:game/:username/:number/:color", play)
	e.POST("/draw/:game/:username", draw)
}

func newGame(c echo.Context) error {

	if err := createNewGame(); err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, &Response{true, newPayload("")}, "  ")
}

func login(c echo.Context) error {
	validGame := joinGame(c.Param("game"), c.Param("username"))
	return respondIfValid(c, (validGame == nil))
}

func startGame(c echo.Context) error {
	dealCards()
	return update(c)
}

func update(c echo.Context) error {
	valid := updateGame(c.Param("game"), c.Param("username"))
	return respondIfValid(c, valid)
}

func cardCount(c echo.Context) error {
	var response *Status
	val, err := getCardCount(c.Param("game"))
	if err != nil {
		response = &Status{"", err.Error()}
	} else {
		response = &Status{val, ""}
	}

	return c.JSONPretty(http.StatusOK, response, "  ")
}

func play(c echo.Context) error {
	// TODO Cards have a value, which can include skip, reverse, etc
	card := model.Card{Color: c.Param("number"), Value: c.Param("color")}
	valid := playCard(c.Param("game"), c.Param("username"), card)
	return respondIfValid(c, valid)
}

func draw(c echo.Context) error {
	valid := drawCard(c.Param("game"), c.Param("username"))
	return respondIfValid(c, valid)
}

func respondIfValid(c echo.Context, valid bool) error {
	var payload *Response
	if valid {
		payload = &Response{true, newPayload(c.Param("username"))}
	} else {
		payload = &Response{false, nil}
	}
	return c.JSONPretty(http.StatusOK, payload, "  ")
}
