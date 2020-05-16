package main

import (
	"net/http"
	"github.com/labstack/echo"
)

// These variables have to be 
type Card struct {
	Number int `json:"number"`
	Color string `json:"color"`
}

type Status struct {
	GameID string `json:"gameID" xml:"gameID"`
	PlayerTurn string `json:"playerTurn" xml:"playerTurn"`
	Cards []Card  `json:"cards"`  // 2 cards for now
}
// User
type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

func setupRoutes(e *echo.Echo) {
	//e.GET("/", hello)
	e.GET("/newgame", newGame)
	e.GET("/startgame", func(c echo.Context) error {
		u := Status {
			GameID: "12234", 
			PlayerTurn: "56123",
			Cards: []Card {
				Card {5,"Blue"},
				Card {3,"Red"},
				Card {6,"Green"},
				Card {8,"Yellow"},
				Card {9,"Red"},
			},
		}
		return c.JSONPretty(http.StatusOK, u, "    ")
	})
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func newGame(c echo.Context) error {
	return c.String(http.StatusOK, "New game")
}
