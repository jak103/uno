package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
)

var sim bool = true

type Response struct {
	ValidGame bool                   `json:"valid"` // Valid game id/game id is in JWT
	Payload   map[string]interface{} `json:"payload"`
}

func setupRoutes(e *echo.Echo) {
	e.GET("/newgame", newGame)
	e.GET("/update/:game/:username", update)
	e.POST("/startgame/:game/:username", startGame)
	e.POST("/login/:game/:username", login)
	e.POST("/play/:game/:username/:number/:color", play)
	e.POST("/draw/:game/:username", draw)
}

func newGame(c echo.Context) error {
	gameid := createNewGame()
	
	// TODO: validate username
	encodedJWT, err := newJWT(c.Param("username"), uuid.New(), gameid, true, []byte(signKey))
	
	payload := newPayload(c.Param("username"), gameid)
	
	if err == nil {
		payload = MakeJWTPayload(payload, encodedJWT)
	} else {
		// TODO: return some sort of error!
		payload = newPayload(c.Param("username"), gameid)
	}
	
	return c.JSONPretty(http.StatusOK, &Response{true, payload}, "  ")
}

func login(c echo.Context) error {
	validGame := joinGame(c.Param("game"), c.Param("username"))
	
	return respondWithJWTIfValid(c, validGame)
}

func startGame(c echo.Context) error {
	dealCards()
	return update(c)
}

func update(c echo.Context) error {
	valid := updateGame(c.Param("game"), c.Param("username"))
	return respondIfValid(c, valid)
}

func play(c echo.Context) error {
	num, _ := strconv.Atoi(c.Param("number"))
	card := Card{num, c.Param("color")}
	valid := playCard(c.Param("game"), c.Param("username"), card)
	return respondIfValid(c, valid)
}

func draw(c echo.Context) error {
	valid := drawCard(c.Param("game"), c.Param("username"))
	return respondIfValid(c, valid)
}

func respondIfValid(c echo.Context, valid bool) error {
	var response *Response
	if valid {
		response = &Response{true, newPayload(c.Param("username"), c.Param("game"))}
	} else {
		response = &Response{false, nil}
	}
	return c.JSONPretty(http.StatusOK, response, "  ")
}

func respondWithJWTIfValid(c echo.Context, valid bool) error {	
	// TODO: validate username and game id
	// TODO: check if they have a JWT before just overriding it! If they do, we need to make a JWT based off of their current one, but add/change the gameid.
	encodedJWT, err := newJWT(c.Param("username"), uuid.New(), c.Param("game"), false, []byte(signKey))
	
	payload := newPayload(c.Param("username"), c.Param("game"))
	
	if err == nil {
		payload = MakeJWTPayload(payload, encodedJWT)
	} else {
		// TODO: return some sort of error!
		
	}
	
	var response *Response
	
	if valid {
		response = &Response{true, payload}
	} else {
		response = &Response{false, nil}
	}
	
	return c.JSONPretty(http.StatusOK, response, "  ")
}
