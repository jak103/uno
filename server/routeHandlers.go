package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/jak103/uno/model"
	"github.com/labstack/echo/v4"
)

var sim bool = true

type Response struct {
	ValidGame bool                   `json:"valid"` // Valid game id
	GameState map[string]interface{} `json:"payload"`
}

func setupRoutes(e *echo.Echo) {
	e.GET("/newgame", newGame)
	e.GET("/update/:game/:username", update)
	e.POST("/startgame/:game/:username", startGame)
	e.POST("/login/:game/:username", login)
	e.POST("/play/:game/:username/:number/:color", play)
	e.POST("/draw/:game/:username", draw)
}

func newGame(c echo.Context) error {
	game, err := createNewGame()

	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, &Response{true, getGameState(game)}, "  ")
}

func login(c echo.Context) error {
	game, err := joinGame(c.Param("game"), c.Param("username"))

	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, &Response{true, getGameState(game)}, "  ")
}

func startGame(c echo.Context) error {
	dealCards(c.Param("game"), c.Param("username"))
	return update(c)
}

func update(c echo.Context) error {
	playerID := getPlayerFromContext(c)
	game, err := updateGame(c.Param("game"), c.Param("username"))
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, &Response{true, getGameState(game, playerID)}, "  ")
}

func play(c echo.Context) error {
	// TODO Cards have a value, which can include skip, reverse, etc
	playerID := getPlayerFromContext(c)
	card := model.Card{c.Param("number"), c.Param("color")}
	game, err := playCard(c.Param("game"), c.Param("username"), card)

	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, &Response{true, getGameState(game, playerID)}, "  ")
}

func draw(c echo.Context) error {
	playerID := getPlayerFromContext(c)
	game, err := drawCard(c.Param("game"), c.Param("username"))

	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, &Response{true, getGameState(game, playerID)}, "  ")

}

func getGameState(game *model.Game, playerID string) map[string]interface{} {
	gameState := make(map[string]interface{})

	// Update known variables
	gameState["direction"] = game.Direction
	gameState["current_player"] = game.CurrentPlayer
	gameState["draw_pile"] = game.DrawPile
	gameState["discard_pile"] = game.DiscardPile
	gameState["game_id"] = game.ID
	gameState["game_over"] = (game.Status == "Finished")

	for _, player := range game.Players {
		if player.ID != playerID {
			for _, card := range player.Cards {
				card.Color = "Blank"
				card.Value = "Blank"
			}
		}
	}

	gameState["all_players"] = game.Players

	return gameState
}

func getPlayerFromContext(c echo.Context) string {
	// TODO Update this to the actual claim key once the JWT team is done
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	playerID := claims["playerID"].(string)

	return playerID
}
