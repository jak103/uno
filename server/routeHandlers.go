package main

import (
	"net/http"
	"github.com/labstack/echo"
	"strconv"
)


var sim bool = true

func setupRoutes(e *echo.Echo) {
	e.GET("/newgame", newGame)
	e.POST("/startgame/:game/:username", startGame)
	e.POST("/login/:game/:username", login)
	e.GET("/update/:game/:username", update)
	e.POST("/play/:game/:username/:number/:color", play)
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
