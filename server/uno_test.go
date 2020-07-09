package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestRandColor(t *testing.T) {

	assert.Equal(t, randColor(0), "red")
	assert.Equal(t, randColor(1), "blue")
	assert.Equal(t, randColor(2), "green")
	assert.Equal(t, randColor(3), "yellow")
	assert.Equal(t, randColor(-1), "")
	assert.Equal(t, randColor(100), "")

}

func TestCreateNewGame(t *testing.T){
	// an attempt to mimic the routehandler's newgame test
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodPost, "/newgame", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	r := &Response{true, newPayload("")}
	r.Payload["game_id"] = "12234"
	assert.Equal(t, checkID("12234"), true)
	assert.Equal(t, createNewGame(c), r)
}

func TestContains(t *testing.T){
	names := []string {"John","Steven","Derek"}

	i, value := contains(names, "Steven")
	assert.Equal(t, 1, i)
	assert.Equal(t, true, value)

	i, value = contains(names, "Jhon")
	assert.Equal(t, -1, i)
	assert.Equal(t, false, value)
}