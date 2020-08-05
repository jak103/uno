package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var unitTestUserName = "UNIT_TEST_USER"
var unitTestGameName = "UNIT_TEST_GAME_NAME"

func TestTotalGamePlayAuth(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)

	// Create an joinable game
	createRec := httptest.NewRecorder()
	createReq := httptest.NewRequest(http.MethodPost, "/api/games?name=game_name&creator=creator_name", nil)
	e.ServeHTTP(createRec, createReq)

	createRsp := make(map[string]interface{})
	json.Unmarshal([]byte(createRec.Body.String()), &createRsp)

	// Create Game Assertions
	assert.NotEqual(t, createRsp["token"], nil)
	assert.NotEqual(t, createRsp["game"], nil)
	game, _ := createRsp["game"].(map[string]interface{})

	// fmt.Println(reflect.TypeOf(game).String())
	// for k, v := range game {
	// 	fmt.Println("k:", k, "v:", v)
	// }

	assert.NotEqual(t, game, nil)
	assert.Equal(t, game["name"], "game_name")
	assert.Equal(t, http.StatusOK, createRec.Code)

	gameID := game["game_id"]

	// fmt.Println("Game ID")
	// fmt.Println(gameID)

	// Test Get Games - all lobby games
	getRec := httptest.NewRecorder()
	getReq := httptest.NewRequest(http.MethodGet, "/api/games", nil)
	e.ServeHTTP(getRec, getReq)

	getRsp := make(map[string]interface{})
	json.Unmarshal([]byte(getRec.Body.String()), &getRsp)

	assert.Equal(t, http.StatusOK, getRec.Code)

	// fmt.Println()
	// fmt.Println()

	// auth
	// config.headers.Authorization = `Token ${token}`;
	// drawReq.Header.Set("Authorization", "Bearer "+loginRes.Payload["JWT"].(string))

	// Test Start Game
	startRec := httptest.NewRecorder()
	startReq := httptest.NewRequest(http.MethodPost, "/api/games/"+gameID.(string)+"/start", nil)
	startReq.Header.Set("Authorization", "Token "+createRsp["token"].(string))

	e.ServeHTTP(startRec, startReq)

	startRsp := make(map[string]interface{})
	json.Unmarshal([]byte(startRec.Body.String()), &startRsp)

	assert.Equal(t, startRsp["name"], "game_name")
	assert.Equal(t, http.StatusOK, startRec.Code)

	// Draw
	drawRec := httptest.NewRecorder()
	drawReq := httptest.NewRequest(http.MethodPost, "/api/games/"+gameID.(string)+"/draw", nil)
	drawReq.Header.Set("Authorization", "Token "+createRsp["token"].(string))

	e.ServeHTTP(drawRec, drawReq)

	drawRsp := make(map[string]interface{})
	json.Unmarshal([]byte(drawRec.Body.String()), &drawRsp)

	assert.Equal(t, drawRsp["name"], "game_name")
	assert.Equal(t, http.StatusOK, drawRec.Code)

	// Play
	playRec := httptest.NewRecorder()
	playReq := httptest.NewRequest(http.MethodPost, "/api/games/"+gameID.(string)+"/draw", nil)
	playReq.Header.Set("Authorization", "Token "+createRsp["token"].(string))

	e.ServeHTTP(playRec, playReq)

	playRsp := make(map[string]interface{})
	json.Unmarshal([]byte(playRec.Body.String()), &playRsp)

	assert.Equal(t, playRsp["name"], "game_name")
	assert.Equal(t, http.StatusOK, playRec.Code)

	// Chat
	chatRec := httptest.NewRecorder()
	chatReq := httptest.NewRequest(http.MethodPost, "/api/chat/"+gameID.(string)+"/add?message=hello_world", nil)
	chatReq.Header.Set("Authorization", "Token "+createRsp["token"].(string))

	e.ServeHTTP(chatRec, chatReq)

	chatRsp := make(map[string]interface{})
	json.Unmarshal([]byte(chatRec.Body.String()), &chatRsp)

	assert.Equal(t, chatRsp["name"], "game_name")
	assert.Equal(t, http.StatusOK, chatRec.Code)

	// Get game state
	stateRec := httptest.NewRecorder()
	stateReq := httptest.NewRequest(http.MethodGet, "/api/games/"+gameID.(string), nil)
	stateReq.Header.Set("Authorization", "Token "+createRsp["token"].(string))

	e.ServeHTTP(stateRec, stateReq)

	stateRsp := make(map[string]interface{})
	json.Unmarshal([]byte(stateRec.Body.String()), &stateRsp)

	assert.Equal(t, stateRsp["name"], "game_name")
	assert.Equal(t, http.StatusOK, stateRec.Code)

	// Test Join Game
	joinRec := httptest.NewRecorder()
	joinReq := httptest.NewRequest(http.MethodPost, "/api/games/"+gameID.(string)+"/join?playerName=player_name", nil)

	e.ServeHTTP(joinRec, joinReq)

	joinRsp := make(map[string]interface{})
	json.Unmarshal([]byte(joinRec.Body.String()), &joinRsp)

	assert.NotEqual(t, joinRsp["token"], nil)
	assert.NotEqual(t, joinRsp["game"], nil)
	assert.Equal(t, http.StatusOK, joinRec.Code)
}
