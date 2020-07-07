package main

import (
	"encoding/json"
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

func TestJoinGame(t *testing.T) {
	e := echo.New()
	setupRoutes(e)

	// new game
	newGameReq := httptest.NewRequest(http.MethodPost, "/newgame", nil)
	newGameReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	newGameRec := httptest.NewRecorder()
	newGameCtx := e.NewContext(newGameReq, newGameRec)

	if assert.NoError(t, newGame(newGameCtx)) {
		assert.Equal(t, http.StatusOK, newGameRec.Code)
		// fmt.Printf("New Game Rsp JSON: %s", newGameRec.Body)
	}
	newGameRsp := Response{}
	json.Unmarshal([]byte(string(newGameRec.Body.Bytes())), &newGameRsp)
	// fmt.Printf("New Game Rsp: %v", newGameRsp)

	assert.Equal(t, newGameRsp.ValidGame, true)

	// login
	var gameID string
	gameID = newGameRsp.Payload["game_id"].(string)
	var username string = "test_user"

	rsp := joinGame(gameID, username)
	assert.Equal(t, rsp.ValidGame, true)
}
