package main

import (
	"net/http"

	"github.com/jak103/uno/model"
	"github.com/labstack/echo/v4"
)

var sim bool = true

type Response struct {
	ValidGame bool                   `json:"valid"` // Valid game id
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
	game, err := createNewGame()

	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, &Response{true, newPayload(game)}, "  ")
}

func login(c echo.Context) error {
	game, err := joinGame(c.Param("game"), c.Param("username"))

	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, &Response{true, newPayload(game)}, "  ")
}

func startGame(c echo.Context) error {
	dealCards(c.Param("game"), c.Param("username"))
	return update(c)
}

func update(c echo.Context) error {
	game, err := updateGame(c.Param("game"), c.Param("username"))
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, &Response{true, newPayload(game)}, "  ")
}

func play(c echo.Context) error {
	// TODO Cards have a value, which can include skip, reverse, etc
	card := model.Card{c.Param("number"), c.Param("color")}
	game, err := playCard(c.Param("game"), c.Param("username"), card)

	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, &Response{true, newPayload(game)}, "  ")
}

func draw(c echo.Context) error {
	game, err := drawCard(c.Param("game"), c.Param("username"))

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
