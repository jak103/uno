package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func setupRoutes(e *echo.Echo) {
	//e.GET("/", hello)
	e.GET("/newgame", newGame)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func newGame(c echo.Context) error {
	return c.String(http.StatusOK, "New game")
}
