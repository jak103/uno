package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func runTest(url string) (*http.Request) {

	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodPost, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	return req
}

func TestNewGame(t *testing.T) {
	req := runTest("/newgame")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	// Assertions
	if assert.NoError(t, newGame(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateGame(t *testing.T) {
	req := runTest("/update/:game/:username")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, update(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestLogin(t *testing.T) {
	req := runTest("/login/:game/:username")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDrawCard(t *testing.T) {
	req := runTest("/draw/:game/:username")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, draw(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}







