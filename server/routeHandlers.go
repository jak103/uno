package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var sim bool = true

type Response struct {
	ValidGame bool                   `json:"valid"` // Valid game id
	Payload   map[string]interface{} `json:"payload"`
}

func setupRoutes(e *echo.Echo) {
	e.GET("/newgame", newGame)
	e.GET("/update", update)
	e.POST("/startgame", startGame)
	e.POST("/login/:game/:username", login)
	e.POST("/play/:number/:color", play)
	e.POST("/draw", draw)
}

func newGame(c echo.Context) error {
	createNewGame()
	return c.JSONPretty(http.StatusOK, &Response{true, newPayload("")}, "  ")
}

func login(c echo.Context) error {
	validGame := joinGame(c.Param("game"), c.Param("username"))
	return respondIfValid(c, validGame, c.Param("username"))
}

func startGame(c echo.Context) error {
	dealCards()
	return update(c)
}

func update(c echo.Context) error {
	claims, validUser := getValidClaims(c.Get(echo.HeaderAuthorization).(string))

	if !validUser {
		return c.JSONPretty(http.StatusUnauthorized, &Response{false, nil}, " ")
	}

	valid := updateGame(claims["gameid"].(string))
	return respondIfValid(c, valid && validUser, claims["userid"].(string))
}

func play(c echo.Context) error {
	claims, validUser := getValidClaims(c.Get(echo.HeaderAuthorization).(string))

	if !validUser {
		return c.JSONPretty(http.StatusUnauthorized, &Response{false, nil}, " ")
	}

	num, _ := strconv.Atoi(c.Param("number"))
	card := Card{num, c.Param("color")}
	valid := playCard(claims["gameid"].(string), claims["userid"].(string), card)
	return respondIfValid(c, valid, claims["userid"].(string))
}

func draw(c echo.Context) error {
	claims, validUser := getValidClaims(c.Get(echo.HeaderAuthorization).(string))

	if !validUser {
		return c.JSONPretty(http.StatusUnauthorized, &Response{false, nil}, " ")
	}

	valid := drawCard(claims["gameid"].(string), claims["userid"].(string))
	return respondIfValid(c, valid, claims["userid"].(string))
}

func respondIfValid(c echo.Context, valid bool, userId string) error {
	var payload *Response
	if valid {
		payload = &Response{true, newPayload(userId)}
	} else {
		payload = &Response{false, nil}
	}
	return c.JSONPretty(http.StatusOK, payload, "  ")
}
