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
