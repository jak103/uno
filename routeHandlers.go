package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func setupRoutes(e *echo.Echo) {
	e.POST("/new_game", newGame)
	e.POST("/:game/join", joinGame)
	e.POST("/:game/start", startGame)
	e.POST("/:game/draw", drawCard)
	e.POST("/:game/play", playCard)
	e.GET("/:game/data", getGameData)
}

func newGame(c echo.Context) error {
	return c.String(http.StatusOK, "New game")
}

func getGameData(c echo.Context) error {
	game := c.Param("game")
	fmt.Println("Game:", game)
	return c.String(http.StatusOK, game + " game data")
}

func joinGame(c echo.Context) error {
	return c.String(http.StatusOK, "Join game")
}

func startGame(c echo.Context) error {
	return c.String(http.StatusOK, "Deal cards")
}

func drawCard(c echo.Context) error {
	return c.String(http.StatusOK, "Draw cards")
}

func playCard(c echo.Context) error {
	return c.String(http.StatusOK, "Play card")
}

/*
/join_game?id={game_id}
====returns====
{'valid': true/false, 'user_id': 12234/-1 }

/new_game?host={username}
====returns====
{'new_id': id, 'player_id': 12234/-1 } of new game

/deal_cards?id={player_id}
====returns====
{'cards' : [ { 'num' : 9, 'clr' : 'red' } , ...] }

/draw?player={player_id}
====returns====
{'cards': [ { 'num': 7, 'clr': 'blue' ] }
*/
