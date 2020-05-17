package main

import (
	"net/http"
	"github.com/labstack/echo"
	// "fmt"
)

var game_id string = "12234"

// These variables have to be 
type Card struct {
	Number int `json:"number"`
	Color string `json:"color"`
}

type Status struct {
	PlayerTurn string `json:"playerTurn" xml:"playerTurn"`
	Cards []Card  `json:"cards"`  // 2 cards for now
}

type Response struct {
	Vaild bool `json:"valid" xml:"valid"`
	Data interface{} `json:"data" xml:"data"`
}

func checkId(c echo.Context) *Response {
	r := &Response{c.Param("id") == "12234", nil}
	return r
}


func setupRoutes(e *echo.Echo) {
	//e.GET("/", hello)
	e.GET("/newgame", newGame)
	e.GET("/startgame/:id", startGame)
	e.POST("/login/:id", login)
}

func login(c echo.Context) error {
	response := checkId(c)
	return c.JSON(http.StatusOK, response)
}

func startGame(c echo.Context) error {
	response := checkId(c)
	response.Data = Status{"56123", []Card {
			Card {5,"Blue"},
			Card {3,"Red"},
			Card {6,"Green"},
			Card {8,"Yellow"},
			Card {9,"Red"},
		},
	}
	return c.JSONPretty(http.StatusOK, response, "    ")
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func newGame(c echo.Context) error {
	return c.String(http.StatusOK, "New game")
}
