package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestTotalGamePlayAuth(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)

	// Login
	loginRec := httptest.NewRecorder()
	loginReq := httptest.NewRequest(http.MethodPost, "/login/tester_name", nil)
	e.ServeHTTP(loginRec, loginReq)
	var loginRes Response
	json.Unmarshal([]byte(loginRec.Body.String()), &loginRes)

	// Login Assertions
	assert.Equal(t, http.StatusOK, loginRec.Code)
	assert.Equal(t, loginRes.ValidGame, true)
	// JWT present
	assert.NotEqual(t, loginRes.Payload["JWT"], nil)

	// New Game
	newGameRec := httptest.NewRecorder()
	newGameReq := httptest.NewRequest(http.MethodGet, "/newgame/", nil)
	newGameReq.Header.Set("Authorization", "Bearer "+loginRes.Payload["JWT"].(string))
	e.ServeHTTP(newGameRec, newGameReq)
	assert.Equal(t, http.StatusOK, newGameRec.Code)

	// Start Game
	// startRec := httptest.NewRecorder()
	// startReq := httptest.NewRequest(http.MethodPost, "/startgame", nil)
	// startReq.Header.Set("Authorization", "Bearer "+loginRes.Payload["JWT"].(string))
	// e.ServeHTTP(startRec, startReq)
	// assert.Equal(t, http.StatusOK, startRec.Code)

	// Draw
	// drawRec := httptest.NewRecorder()
	// drawReq := httptest.NewRequest(http.MethodPost, "/draw", nil)
	// drawReq.Header.Set("Authorization", "Bearer "+loginRes.Payload["JWT"].(string))
	// drawReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// e.ServeHTTP(drawRec, drawReq)
	// assert.Equal(t, http.StatusOK, drawRec.Code)

	// Play
	// playRec := httptest.NewRecorder()
	// playReq := httptest.NewRequest(http.MethodPost, "/play/1/blue", nil)
	// playReq.Header.Set("Authorization", "Bearer "+loginRes.Payload["JWT"].(string))
	// e.ServeHTTP(playRec, playReq)

	// assert.Equal(t, http.StatusOK, playRec.Code)

	// Update
	// updateRec := httptest.NewRecorder()
	// updateReq := httptest.NewRequest(http.MethodGet, "/update", nil)
	// updateReq.Header.Set("Authorization", "Bearer "+loginRes.Payload["JWT"].(string))
	// e.ServeHTTP(updateRec, updateReq)
	// assert.Equal(t, http.StatusOK, updateRec.Code)
}

func TestLogin_error(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	assert.Error(t, login(c)) // This should error because no game was created
}

func TestDraw_error(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodPost, "/draw", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	assert.Error(t, login(c)) // This should error because no game was created
}

func TestUpdate_error(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodGet, "/update", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	assert.Error(t, login(c)) // This should error because no game was created
}

func TestPlay_error(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodGet, "/play", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// TODO: finish mocking up an auth header here, and in other tests
	/*isHost := true
	  encodedJWT, err := newJWT("Thomas", "userid", "gameid", isHost, []byte(signKey))

	  assert.Equal(t, nil, err)

	  req.Header.Set(echo.HeaderAuthorization, "bearer" + encodedJWT)*/

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	assert.Error(t, login(c)) // This should error because no game was created
}

func TestStartGame_error(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	req := httptest.NewRequest(http.MethodGet, "/startgame", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	assert.Error(t, login(c)) // This should error because no game was created
}
