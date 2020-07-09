package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func runTest(url, urlMethod) {

	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodPost, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, urlMethod(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func NewGame(t *testing.T) {
	runTest("/newgame", newGame())
}

func TestUpdate(t *testing.T) {
	runTest("/update/:game/:username", update())
}

func TestLogin(t *testing.T) {
	runTest("/login/:game/:username", login())
}

func TestDraw(t *testing.T) {
	runTest("/draw/:game/:username", draw())
}







