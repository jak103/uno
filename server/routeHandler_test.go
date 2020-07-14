package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodPost, "/newgame", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, newGame(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdate(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodPost, "/update/:12234/:testUser", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, update(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestLogin(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodPost, "/login/:12234/:testUser", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestPlay(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodPost, "/play/:12234/:testUser/:1/:red", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, play(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDraw(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodPost, "/draw/:12234/:testUser", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, draw(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCheckForWinner(t *testing.T) {
	players = []string{"testUser"}
	allCards = make(map[string][]Card)
	dealCards()
	assert.Equal(t, "", checkForWinner())
	allCards[players[0]] = make([]Card, 0)
	assert.Equal(t, "testUser", checkForWinner())
}

func TestStartGame(t *testing.T) {
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodPost, "/startgame/12234/testUser", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
