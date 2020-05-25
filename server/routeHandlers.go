package main

import (
	"net/http"
	"github.com/labstack/echo"
)


var sim bool = true

func setupRoutes(e *echo.Echo) {
	e.GET("/newgame", newGame)
	e.POST("/startgame/:game/:username", startGame)
	e.POST("/login/:game/:username", login)
	e.GET("/update/:game/:username", update)
}

func newGame(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, createNewGame(c), "    ")
}

func login(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, joinGame(c), "    ")
}

func startGame(c echo.Context) error {
	dealCards()
	return c.JSONPretty(http.StatusOK, update(c), "    ")
}

func update(c echo.Context) error {
	// if sim {
	// 	resp := (c)

	// } else {
	return c.JSONPretty(http.StatusOK, updateGame(c), "    ")
	// }
}
