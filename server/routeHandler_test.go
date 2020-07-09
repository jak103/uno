package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

// Route handlers unit tests

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

func TestLogin(t *testing.T) {
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdate(t *testing.T) {
    e := echo.New()
    setupRoutes(e)
    req := httptest.NewRequest(http.MethodPost, "/update", nil)
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    if assert.NoError(t, update(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)
    }
}