package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func createMockServerAndRequest() (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodPost, "/newgame", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return c, rec
}

//TestRespondIfValid this wont bring the code coverage up but its worth it to test this if it breaks on this level.
func TestRespondIfValid(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodGet, "/respondIfValid", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, respondIfValid(c, true)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestNewGame(t *testing.T) {
	// Setup
	c, rec := createMockServerAndRequest()

	// Assertions
	if assert.NoError(t, newGame(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// store response in map and make sure valid field is true
		var recData map[string]interface{}
		json.Unmarshal([]byte(rec.Body.String()), &recData)
		assert.Equal(t, true, recData["valid"])
	}
}

func TestLogin(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDraw(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodPost, "/draw", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, draw(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdate(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodGet, "/update", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, update(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestPlay(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodGet, "/play", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, play(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestStartGame(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodGet, "/startgame", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	players = []string{"player1", "player2"}
	// Assertions
	if assert.NoError(t, startGame(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
