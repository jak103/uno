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
		var rsp Response
		json.Unmarshal([]byte(rec.Body.String()), &rsp)
		assert.Equal(t, true, rsp.ValidGame)
		assert.Equal(t, "12234", rsp.Payload["game_id"])
	}
}

func TestLogin(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	newGameReq := httptest.NewRequest(http.MethodPost, "/newgame", nil)
	newGameReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	newGameRec := httptest.NewRecorder()
	c := e.NewContext(newGameReq, newGameRec)

	// Assertions
	if assert.NoError(t, newGame(c)) {
		assert.Equal(t, http.StatusOK, newGameRec.Code)
		var rsp Response
		json.Unmarshal([]byte(newGameRec.Body.String()), &rsp)
		assert.Equal(t, true, rsp.ValidGame)
		assert.Equal(t, "12234", rsp.Payload["game_id"])
	}

	loginReq := httptest.NewRequest(http.MethodPost, "/login/:game/:username", nil)
	loginRec := httptest.NewRecorder()
	c1 := e.NewContext(loginReq, loginRec)
	c1.SetParamNames("game", "username")
	c1.SetParamValues("911", "test")

	if assert.NoError(t, login(c1)) {
		assert.Equal(t, http.StatusOK, loginRec.Code)
		var rsp Response
		json.Unmarshal([]byte(loginRec.Body.String()), &rsp)
		assert.Equal(t, false, rsp.ValidGame)
	}

	loginReq2 := httptest.NewRequest(http.MethodPost, "/login/:game/:username", nil)
	loginRec2 := httptest.NewRecorder()
	c2 := e.NewContext(loginReq2, loginRec2)
	c2.SetParamNames("game", "username")
	c2.SetParamValues("12234", "test")

	if assert.NoError(t, login(c2)) {
		assert.Equal(t, http.StatusOK, loginRec2.Code)
		var rsp Response
		json.Unmarshal([]byte(loginRec2.Body.String()), &rsp)
		assert.Equal(t, true, rsp.ValidGame)
	}
}

func TestStartGame(t *testing.T) {
	// Setup
	e := echo.New()
	setupRoutes(e)
	newGameReq := httptest.NewRequest(http.MethodPost, "/newgame", nil)
	newGameReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	newGameRec := httptest.NewRecorder()
	c := e.NewContext(newGameReq, newGameRec)

	// Assertions
	if assert.NoError(t, newGame(c)) {
		assert.Equal(t, http.StatusOK, newGameRec.Code)
		var rsp Response
		json.Unmarshal([]byte(newGameRec.Body.String()), &rsp)
		assert.Equal(t, true, rsp.ValidGame)
		assert.Equal(t, "12234", rsp.Payload["game_id"])
	}

	loginReq := httptest.NewRequest(http.MethodPost, "/login/:game/:username", nil)
	loginRec := httptest.NewRecorder()
	c1 := e.NewContext(loginReq, loginRec)
	c1.SetParamNames("game", "username")
	c1.SetParamValues("12234", "unoTester")

	if assert.NoError(t, login(c1)) {
		assert.Equal(t, http.StatusOK, loginRec.Code)
		var rsp Response
		json.Unmarshal([]byte(loginRec.Body.String()), &rsp)
		assert.Equal(t, true, rsp.ValidGame)
	}

	startReq := httptest.NewRequest(http.MethodPost, "/startgame/:game/:username", nil)
	startRec := httptest.NewRecorder()
	c3 := e.NewContext(startReq, startRec)
	c3.SetParamNames("game", "username")
	c3.SetParamValues("12234", "unoTester")

	if assert.NoError(t, startGame(c3)) {
		assert.Equal(t, http.StatusOK, startRec.Code)
		var rsp Response
		json.Unmarshal([]byte(startRec.Body.String()), &rsp)
		assert.Equal(t, true, rsp.ValidGame)
		assert.Equal(t, "", rsp.Payload["game_over"])
	}
}
