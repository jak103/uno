package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func createEcho() *echo.Echo {
	e := echo.New()
	setupRoutes(e)
	return e
}

func createRequest(e *echo.Echo, method, path string) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, path, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

func TestNewGame(t *testing.T) {
	fmt.Println("New Game")
	// Setup
	e := echo.New()
	setupRoutes(e)
	c, rec := createRequest(e, http.MethodGet, "/newgame")

	// Assertions
	if assert.NoError(t, newGame(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		fmt.Println(rec.Body)
	}
}

func TestLogin(t *testing.T) {
	fmt.Println("Join Game")
	// Setup
	e := echo.New()
	setupRoutes(e)

	contextA, _ := createRequest(e, http.MethodGet, "/newgame")
	newGame(contextA)

	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorderB := httptest.NewRecorder()
	contextB := e.NewContext(req, recorderB)
	contextB.SetParamNames("game", "username")
	contextB.SetParamValues("12234", "BudderBoy")

	// Assertions
	if assert.NoError(t, login(contextB)) {
		assert.Equal(t, http.StatusOK, recorderB.Code)
		fmt.Println(recorderB.Body)
	}
}

func TestStartGame(t *testing.T) {
	fmt.Println("Start Game")
	// Setup
	e := createEcho()

	// Create Game
	c, rec := createRequest(e, http.MethodGet, "/newgame")
	newGame(c)

	// Join Game
	c, rec = createRequest(e, http.MethodPost, "/login")
	c.SetParamNames("game", "username")
	c.SetParamValues("12234", "BudderBoy")
	login(c)

	//Start Game
	c, rec = createRequest(e, http.MethodPost, "/startgame")
	c.SetParamNames("game", "username")
	c.SetParamValues("12234", "BudderBoy")
	// Assertions
	if assert.NoError(t, startGame(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		fmt.Println(rec.Body)
	}
}

func TestPlay(t *testing.T) {
	fmt.Println("Test Play")
	// Setup
	e := createEcho()

	// Create Game
	c, rec := createRequest(e, http.MethodGet, "/newgame")
	newGame(c)

	// Join Game
	c, rec = createRequest(e, http.MethodPost, "/login")
	c.SetParamNames("game", "username")
	c.SetParamValues("12234", "BudderBoy")
	login(c)

	//Start Game
	c, rec = createRequest(e, http.MethodPost, "/startgame")
	c.SetParamNames("game", "username")
	c.SetParamValues("12234", "BudderBoy")
	play(c)

	//Play
	c, rec = createRequest(e, http.MethodPost, "/play")
	c.SetParamNames("game", "username", "number", "color")
	c.SetParamValues("12234", "BudderBoy", "1", "red")

	// Assertions
	if assert.NoError(t, play(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		fmt.Println(rec.Body)
	}
}
