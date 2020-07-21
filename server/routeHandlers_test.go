package main

import (
  "encoding/json"
  "net/http/httptest"
  "net/http"
  "testing"

  "github.com/labstack/echo"
  "github.com/stretchr/testify/assert"
)

// TODO: Is there a way to compose these tests? We are making a LOT of the same calls...

// Create a mock server
func createMockServer() *echo.Echo {
  e := echo.New()
  setupRoutes(e)
  return e
}

// Test creation of a new game. 
func TestNewGame(t *testing.T) {
  e := createMockServer()

  // Set up the mock request for this test and capture response
  req := httptest.NewRequest(http.MethodGet, "/newgame", nil)
  req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
  rec := httptest.NewRecorder()
  c := e.NewContext(req, rec)

  // Pass the mock request through our route handlers
  rsp := newGame(c)

  if assert.NoError(t, rsp) {
    // Make sure we got an OK return code
    assert.Equal(t, http.StatusOK, rec.Code)

    // Read the JSON response into a map
    var jsonBody map[string]interface{}
    json.Unmarshal([]byte(rec.Body.String()), &jsonBody)

    // Check that "valid" value is true
    assert.Equal(t, true, jsonBody["valid"])
  }
}

// Test logging into a game after creating it
func TestLogin(t *testing.T) {
  e := createMockServer()

  var gameID string
  username := "testUser"
  
  // Set up the mock request to create a new game. We will capture the gameID from this request
  {
    req := httptest.NewRequest(http.MethodGet, "/newgame", nil)
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    // Pass the mock request through our route handlers
    rsp := newGame(c)

    if assert.NoError(t, rsp) {
      // Make sure we got an OK return code
      assert.Equal(t, rec.Code, http.StatusOK)

      // Read the JSON response into a map
      var jsonBody map[string]interface{}
      json.Unmarshal([]byte(rec.Body.String()), &jsonBody)

      // Check that "valid" value is true
      assert.Equal(t, true, jsonBody["valid"])
      
      gameID = jsonBody["payload"].(map[string]interface {})["game_id"].(string)
    }
  }

  {
    // Now that we have the game ID, lets try starting it
    req := httptest.NewRequest(http.MethodPost, "/", nil)
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    // TRICKY: We have to set the path parameters manually using these functions for testing
    c.SetPath("/login/:game/:username")
    c.SetParamNames("game", "username")
    c.SetParamValues(gameID, username)

    // Pass the mock request through our route handlers
    rsp := login(c)

    if assert.NoError(t, rsp) {
      // Make sure we got an OK return code
      assert.Equal(t, http.StatusOK, rec.Code)

      // Read the JSON response into a map
      var jsonBody map[string]interface{}
      json.Unmarshal([]byte(rec.Body.String()), &jsonBody)

      // Check that "valid" value is true
      assert.Equal(t, true, jsonBody["valid"])
    }
  }
}

// Test starting a game after logging into it
func TestStartGame(t *testing.T) {
  e := createMockServer()

  var gameID string
  username := "testUser"
  
  // Set up the mock request to create a new game. We will capture the gameID from this request
  {
    req := httptest.NewRequest(http.MethodGet, "/newgame", nil)
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    // Pass the mock request through our route handlers
    rsp := newGame(c)

    if assert.NoError(t, rsp) {
      // Make sure we got an OK return code
      assert.Equal(t, http.StatusOK, rec.Code)

      // Read the JSON response into a map
      var jsonBody map[string]interface{}
      json.Unmarshal([]byte(rec.Body.String()), &jsonBody)

      // Check that "valid" value is true
      assert.Equal(t, jsonBody["valid"], true)
      
      gameID = jsonBody["payload"].(map[string]interface {})["game_id"].(string)
    }
  }

  {
    // Now that we have the game ID, lets try starting it
    req := httptest.NewRequest(http.MethodPost, "/", nil)
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    // TRICKY: We have to set the path parameters manually using these functions for testing
    c.SetPath("/login/:game/:username")
    c.SetParamNames("game", "username")
    c.SetParamValues(gameID, username)

    // Pass the mock request through our route handlers
    rsp := login(c)

    if assert.NoError(t, rsp) {
      // Make sure we got an OK return code
      assert.Equal(t, http.StatusOK, rec.Code)

      // Read the JSON response into a map
      var jsonBody map[string]interface{}
      json.Unmarshal([]byte(rec.Body.String()), &jsonBody)

      // Check that "valid" value is true
      assert.Equal(t, jsonBody["valid"], true)
    }
  }

  // Now that we are logged in, start the game
  {
    // Now that we have the game ID, lets try starting it
    req := httptest.NewRequest(http.MethodPost, "/", nil)
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    // TRICKY: We have to set the path parameters manually using these functions for testing
    c.SetPath("/startgame/:game/:username")
    c.SetParamNames("game", "username")
    c.SetParamValues(gameID, username)

    // Pass the mock request through our route handlers
    rsp := startGame(c)

    if assert.NoError(t, rsp) {
      // Make sure we got an OK return code
      assert.Equal(t, http.StatusOK, rec.Code)

      // We do not expect a resposne body, just the code.
    }
  }
}
