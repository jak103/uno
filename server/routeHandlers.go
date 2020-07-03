package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

var sim bool = true

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
	return c.JSONPretty(http.StatusOK, joinGame(c), "  ")
}

func startGame(c echo.Context) error {
	dealCards()
	return c.JSONPretty(http.StatusOK, update(c), "  ")
}

func update(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, updateGame(c), "  ")
}

func play(c echo.Context) error {
	num, _ := strconv.Atoi(c.Param("number"))

	return c.JSONPretty(http.StatusOK, playCard(c, Card{num, c.Param("color")}), "  ")
}

func draw(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, drawCard(c), "  ")
}
