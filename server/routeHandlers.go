package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func setupRoutes(e *echo.Echo) {
	e.GET("/newgame", newGame)
	e.GET("/login/:game/:username", login)
	e.GET("/startgame/:game/:username", startGame)
	// e.POST("/startgame/:game/:username", startGame)
	// e.POST("/login/:game/:username", login)
	e.GET("/update/:game/:username", update)
}

func newGame(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, createNewGame(c), "    ")
}

func login(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, joinGame(c), "    ")
}

func startGame(c echo.Context) error {

	response := newResponse(c)
	dealCards(response)

	return c.JSONPretty(http.StatusOK, response, "    ")
}

func update(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, newResponse(c), "    ")
}
