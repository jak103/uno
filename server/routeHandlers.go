package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func setupRoutes(e *echo.Echo) {
	e.GET("/newgame", newGame)
	e.GET("/update/:game/:username", update)
	e.POST("/startgame/:game/:username", startGame)
	e.POST("/login/:game/:username", login)
	e.POST("/play/:game/:username/:number/:color", play)
	e.POST("/draw/:game/:username", draw)
}

func newGame(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, createNewGame(c), "  ")
}

func login(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, joinGameIfValid(c), "  ")
}

func startGame(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, dealCardsAndStartGame(c), "  ")
}

func update(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, heartBeat(c), "  ")
}

func play(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, playCard(c), "  ")
}

func draw(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, drawCard(c), "  ")
}
