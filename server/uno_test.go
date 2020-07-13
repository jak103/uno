package main

import (
	"net/http"
	"net/http/httptest"

	"encoding/json"

	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func parse(rec *httptest.ResponseRecorder) Response {
	var res Response
	json.Unmarshal([]byte(rec.Body.String()), &res)
	return res
}

func serve(e *echo.Echo, method string, url string) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, httptest.NewRequest(method, url, nil))
	return rec
}

func TestLogin(t *testing.T) {
	e := echo.New()

	// Setup server routes
	setupRoutes(e)

	// Ensure logging into an existing game fails properly
	rec := serve(e, http.MethodPost, "/login/0/Tyler")
	assert.Equal(t, http.StatusOK, rec.Code)
	res := parse(rec)
	assert.Equal(t, false, res.ValidGame)
	assert.Equal(t, 0, len(res.Payload))

	// Create a new game
	game := parse(serve(e, http.MethodGet, "/newgame")).Payload["game_id"].(string)

	// Login as "Tyler" and ensure a new player and cards are added
	rec = serve(e, http.MethodPost, "/login/" + game + "/Tyler")
	assert.Equal(t, http.StatusOK, rec.Code)
	res = parse(rec)
	assert.Equal(t, true, res.ValidGame)
	players := res.Payload["all_players"].([]interface{})
	assert.Equal(t, 1, len(players))
	assert.Equal(t, "Tyler", players[0].(string))

	// Login as "Tyler" again and ensure a duplicate player isn't added
	rec = serve(e, http.MethodPost, "/login/" + game + "/Tyler")
	assert.Equal(t, http.StatusOK, rec.Code)
	res = parse(rec)
	assert.Equal(t, true, res.ValidGame)
	players = res.Payload["all_players"].([]interface{})
	assert.Equal(t, 1, len(players))
	assert.Equal(t, "Tyler", players[0].(string))

	// Start the game
	serve(e, http.MethodPost, "/startgame/" + game + "/Tyler")

	// Login as "Celeste" and ensure it fails
/*
	// At the time of writing, the world was not ready for such greatness as this test.
	rec = serve(e, http.MethodPost, "/login/" + game + "/Tyler")
	assert.Equal(t, http.StatusOK, rec.Code)
	res = parse(rec)
	assert.Equal(t, false, res.ValidGame)
	assert.Equal(t, 0, len(res.Payload))
*/
}
