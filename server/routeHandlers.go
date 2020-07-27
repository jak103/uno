package main

import (
	"net/http"
	"github.com/google/uuid"
	"github.com/jak103/uno/model"
	"github.com/labstack/echo/v4"
)

var sim bool = true

type Response struct {
	ValidGame bool                   `json:"valid"` // Valid game id/game id is in JWT
	Payload   map[string]interface{} `json:"payload"`
}

func setupRoutes(e *echo.Echo) {
	e.GET("/newgame/:username", newGame)
	e.GET("/update", update)
	e.POST("/startgame", startGame)
	e.POST("/login/:game/:username", login)
	e.POST("/play/:number/:color", play)
	e.POST("/draw", draw)
}

func newGame(c echo.Context) error {
	gameid, gameErr := createNewGame()
    
    if  gameErr != nil {
		return gameErr
	}
    
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
    err := joinGame(c.Param("game"), c.Param("username"))
    
    return respondWithJWTIfValid(c, err == nil)
}

func startGame(c echo.Context) error {
	dealCards()
	return update(c)
}

func update(c echo.Context) error {
    
    authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
    if authHeader == "" {
        return c.JSONPretty(http.StatusUnauthorized, &Response{false, nil}, " ")
    }
	claims, validUser := getValidClaims(authHeader)

	if !validUser {
		return c.JSONPretty(http.StatusUnauthorized, &Response{false, nil}, " ")
	}

	valid := updateGame(claims["gameid"].(string))
	return respondIfValid(c, valid && validUser, claims["userid"].(string), claims["gameid"].(string))
}

func play(c echo.Context) error {
    
    authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
    if authHeader == "" {
        return c.JSONPretty(http.StatusUnauthorized, &Response{false, nil}, " ")
    }
	claims, validUser := getValidClaims(authHeader)

	if !validUser {
		return c.JSONPretty(http.StatusUnauthorized, &Response{false, nil}, " ")
	}
    
    // TODO Cards have a value, which can include skip, reverse, etc
    card := model.Card{c.Param("number"), c.Param("color")}
	valid := playCard(claims["gameid"].(string), claims["userid"].(string), card)
	return respondIfValid(c, valid, claims["userid"].(string), claims["gameid"].(string))
}

func draw(c echo.Context) error {
    
    authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
    if authHeader == "" {
        return c.JSONPretty(http.StatusUnauthorized, &Response{false, nil}, " ")
    }
	claims, validUser := getValidClaims(authHeader)

	if !validUser {
		return c.JSONPretty(http.StatusUnauthorized, &Response{false, nil}, " ")
	}

	valid := drawCard(claims["gameid"].(string), claims["userid"].(string))
	return respondIfValid(c, valid, claims["userid"].(string), claims["gameid"].(string))
}

func respondIfValid(c echo.Context, valid bool, userId string, gameId string) error {
	var response *Response
	if valid {
		response = &Response{true, newPayload(userId, gameId)}
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
