package main

import (
	"net/http"
    "github.com/jak103/uno/db"
	"github.com/jak103/uno/model"
	"github.com/labstack/echo/v4"
)

var sim bool = true

type Response struct {
	ValidGame bool                   `json:"valid"` // Valid game id/game id is in JWT
	Payload   map[string]interface{} `json:"payload"`
}

func setupRoutes(e *echo.Echo) {
	e.GET("/newgame/", newGame)
	e.GET("/update", update)
	e.POST("/startgame", startGame)
	e.POST("/login/:username", login)
    e.POST("/joinGame/:game", join)
	e.POST("/play/:number/:color", play)
	e.POST("/draw", draw)
}

func newGame(c echo.Context) error {
	game, gameErr := createNewGame()
    
    if  gameErr != nil {
		return gameErr
	}
    
    return c.JSONPretty(http.StatusOK, &Response{true, newPayload(game)}, "  ")
}

func login(c echo.Context) error {
	username := c.Param("username")
    
    database, err := db.GetDb()
	if err != nil {
		return err
	}
    
    player, playerErr := database.CreatePlayer(username)

	if playerErr != nil {
		return playerErr
	}
    
    token, err := newJWT(username, player.ID);
    
    if err != nil {
		return err
	}
    
	return c.JSONPretty(http.StatusOK, &Response{true, makeJWTPayload(token)}, "  ")
}

func join(c echo.Context) error {
    authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
    player, validPlayer, err := getPlayerFromHeader(authHeader)
    
    if err != nil {
		return err
	}
    
    if !validPlayer {
        return c.JSONPretty(http.StatusUnauthorized, &Response{false, nil}, " ")
    }
    
	game, err := joinGame(c.Param("game"), player)
    
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, &Response{true, newPayload(game)}, "  ")
}

func startGame(c echo.Context) error {
    authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
    player, validPlayer, err := getPlayerFromHeader(authHeader)
    
    if err != nil {
		return err
	}
    
    if !validPlayer {
        return c.JSONPretty(http.StatusUnauthorized, &Response{false, nil}, " ")
    }
    
	dealCards(c.Param("game"), player)
	return update(c)
}

func update(c echo.Context) error {
    authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
    player, validPlayer, err := getPlayerFromHeader(authHeader)
    
    if err != nil {
		return err
	}
    
    if !validPlayer {
        return c.JSONPretty(http.StatusUnauthorized, &Response{false, nil}, " ")
    }
    
	game, err := updateGame(c.Param("game"), player)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, &Response{true, newPayload(game)}, "  ")
}

func play(c echo.Context) error {
	// TODO Cards have a value, which can include skip, reverse, etc
	card := model.Card{c.Param("number"), c.Param("color")}
    
    authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
    player, validPlayer, err := getPlayerFromHeader(authHeader)
    
    if err != nil {
		return err
	}
    
    if !validPlayer {
        return c.JSONPretty(http.StatusUnauthorized, &Response{false, nil}, " ")
    }
    
	game, err := playCard(c.Param("game"), player, card)

	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, &Response{true, newPayload(game)}, "  ")
}

func draw(c echo.Context) error {
    authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
    player, validPlayer, err := getPlayerFromHeader(authHeader)
    
    if err != nil {
		return err
	}
    
    if !validPlayer {
        return c.JSONPretty(http.StatusUnauthorized, &Response{false, nil}, " ")
    }

	game, err := drawCard(c.Param("game"), player)

	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, &Response{true, newPayload(game)}, "  ")

}

func newPayload(game *model.Game) map[string]interface{} {
	payload := make(map[string]interface{})

	// Update known variables
	payload["direction"] = game.Direction
	payload["current_player"] = game.CurrentPlayer
	payload["all_players"] = game.Players
	payload["draw_pile"] = game.DrawPile
	payload["discard_pile"] = game.DiscardPile
	payload["game_id"] = game.ID
	payload["game_over"] = game.Status == "Finished"

	return payload
}
