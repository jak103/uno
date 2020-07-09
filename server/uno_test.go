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

// func TestPlayCard(t *testing.T){
// 	e := echo.New()
// 	setupRoutes(e)
// 	gameID = "12234"
// 	currCard[0] = Card{1,"red"}
	
// 	req := httptest.NewRequest(http.MethodPost, "/play/12234/me/3/red", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
	
// 	r := playCard(c,Card{1,randColor(3)})
// 	assert.Equal(t,r.ValidGame, false)

// }