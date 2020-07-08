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

//	"http://localhost:8080/play/" +
//	this.game_id +
//	"/" +
//	this.username +
//	"/" +
//	card.number +
//	"/" +
//	card.color

func ValidCardTest(t *testing.T) {

	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodPost, "/newgame", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, )
}

// Make sure that the Username is Set
// Make sure that the Usernames accross computers are not overlapping
func UserNameTest(t *testing.T) {

	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodPost, "/newgame", nil)

	if assert.NoError(t, )
}

func YourTurnTest(t *testing.T) {


}

func WonTheGameTest(t *testing.T) {


}
