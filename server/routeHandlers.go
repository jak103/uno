package main

import (
	"net/http"
	"github.com/labstack/echo"
	)

var game_id string = "12234"

type Card struct {
	Number int `json:"number"`
	Color string `json:"color"`
}

type Response struct {
	Valid bool `json:"valid"`  // Valid game id
	Payload map[string]interface{} `json:"payload"`
}

func checkId(c echo.Context) *Response {
	r := &Response{c.Param("id") == "12234", nil}
	return r
}

func newMap() map[string]interface{} {
	tmp := make(map[string]interface{})
	tmp["current_card"] = []*Card{&Card{4, "red"}}
	return tmp
}


func setupRoutes(e *echo.Echo) {
	//e.GET("/", hello)
	e.GET("/newgame", newGame)
	e.POST("/startgame/:id", startGame)
	e.POST("/login/:id", login)
	e.POST("/newgame", newGame)
}


func login(c echo.Context) error {
	response := checkId(c)
	return c.JSON(http.StatusOK, response)
}

func startGame(c echo.Context) error {
	response := checkId(c)
	pay := newMap()

	pay["cards"] = []*Card {
			&Card{5,"blue"},
			&Card{3,"red"},
			&Card{6,"green"},
			&Card{8,"yellow"},
			&Card{9,"red"},
	}
	pay["current_player"] = "Ryan"
	pay["players"] = []string{"Bill", "Bob", "Jill", "Ryan"}
	
	response.Payload = pay
	return c.JSON(http.StatusOK, response)
}


func newGame(c echo.Context) error {
	pay := newMap()
	pay["game_id"] = 12234
	return c.JSON(http.StatusOK, &Response{true, pay})
}
