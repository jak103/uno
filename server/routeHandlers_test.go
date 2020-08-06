package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/jak103/uno/db"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestTotalGamePlayAuth(t *testing.T) {
	// Inital Setup
	e := echo.New()
	setupRoutes(e)

	// initalize a database
	db.GetDb()

	// Create an joinable game
	createRec := httptest.NewRecorder()
	createReq := httptest.NewRequest(http.MethodPost, "/api/games?name=game_name&creator=creator_name", nil)
	createCtx := e.NewContext(createReq, createRec)

	var gameID string
	var token *jwt.Token
	var validToken bool
	if assert.NoError(t, newGame(createCtx)) {
		createRsp := make(map[string]interface{})
		json.Unmarshal([]byte(createRec.Body.String()), &createRsp)

		// Create Game Assertions
		assert.NotEqual(t, createRsp["token"], nil)

		token, validToken = parseJWT(createRsp["token"].(string), tokenSecret)

		assert.Equal(t, true, validToken)

		assert.NotEqual(t, createRsp["game"], nil)
		game, _ := createRsp["game"].(map[string]interface{})

		assert.NotEqual(t, game, nil)
		assert.Equal(t, game["name"], "game_name")
		assert.Equal(t, http.StatusOK, createRec.Code)

		gameID = game["game_id"].(string)
	}

	// Test Get Games - all lobby games
	getRec := httptest.NewRecorder()
	getReq := httptest.NewRequest(http.MethodGet, "/api/games", nil)
	getCtx := e.NewContext(getReq, getRec)

	if assert.NoError(t, getGames(getCtx)) {
		assert.Equal(t, http.StatusOK, getRec.Code)
	}

	// Test Start Game
	startRec := httptest.NewRecorder()
	startReq := httptest.NewRequest(http.MethodPost, "/api/games/"+gameID+"/start", nil)

	startCtx := e.NewContext(startReq, startRec)
	startCtx.Set("user", token)
	startCtx.SetParamNames("id")
	startCtx.SetParamValues(gameID)

	if assert.NoError(t, startGame(startCtx)) {
		assert.Equal(t, http.StatusOK, startRec.Code)
	}

	// Draw
	drawRec := httptest.NewRecorder()
	drawReq := httptest.NewRequest(http.MethodPost, "/api/games/"+gameID+"/draw", nil)

	drawCtx := e.NewContext(drawReq, drawRec)
	drawCtx.Set("user", token)
	drawCtx.SetParamNames("id")
	drawCtx.SetParamValues(gameID)

	if assert.NoError(t, draw(drawCtx)) {
		assert.Equal(t, http.StatusOK, drawRec.Code)
	}

	// Play
	playRec := httptest.NewRecorder()
	playReq := httptest.NewRequest(http.MethodPost, "/api/games/"+gameID+"/draw", nil)

	playCtx := e.NewContext(playReq, playRec)
	playCtx.Set("user", token)
	playCtx.SetParamNames("id")
	playCtx.SetParamValues(gameID)

	if assert.NoError(t, play(playCtx)) {
		assert.Equal(t, http.StatusOK, playRec.Code)
	}

	// Chat
	chatRec := httptest.NewRecorder()
	chatReq := httptest.NewRequest(http.MethodPost, "/api/chat/"+gameID+"/add?playerName=creator_name&message=hello_world", nil)

	chatCtx := e.NewContext(chatReq, chatRec)
	chatCtx.Set("user", token)
	chatCtx.SetParamNames("id")
	chatCtx.SetParamValues(gameID)

	if assert.NoError(t, addNewMessage(chatCtx)) {
		assert.Equal(t, http.StatusOK, chatRec.Code)
	}

	// Get game state
	stateRec := httptest.NewRecorder()
	stateReq := httptest.NewRequest(http.MethodGet, "/api/games/"+gameID, nil)

	e.NewContext(stateReq, stateRec)
	stateCtx := e.NewContext(stateReq, stateRec)
	stateCtx.Set("user", token)
	stateCtx.SetParamNames("id")
	stateCtx.SetParamValues(gameID)

	if assert.NoError(t, getGameState(stateCtx)) {
		assert.Equal(t, http.StatusOK, stateRec.Code)
	}

	// Test Join Game
	joinRec := httptest.NewRecorder()
	joinReq := httptest.NewRequest(http.MethodPost, "/api/games/"+gameID+"/join?playerName=player_name", nil)

	joinCtx := e.NewContext(joinReq, joinRec)
	joinCtx.SetParamNames("id")
	joinCtx.SetParamValues(gameID)
	joinCtx.Set("playerName", "player_name")

	if assert.NoError(t, joinExistingGame(joinCtx)) {
		assert.Equal(t, http.StatusOK, joinRec.Code)
	}
}
