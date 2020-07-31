package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/jak103/uno/db"
	"github.com/jak103/uno/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var sim bool = true

func setupRoutes(e *echo.Echo) {
	// Routes that don't require a valid JWT
	e.GET("/games", getGames)
	e.POST("/games", newGame)

	// Create a group that requires a valid JWT
	group := e.Group("/api")

	group.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("usudevops"),
		AuthScheme: "Token",
	}))
	/*
		group.POST("/games/:id/join", joinGame)
		group.POST("/games/:id/start", startGame)
		group.POST("/games/:id/play", play)
		group.POST("/games/:id/draw", draw)
		group.POST("/games/:id/uno", callUno)

		group.GET("/games/:id", getGameState)
	*/
}

func getGames(c echo.Context) error {
	//log.Println("Running getGames")
	database, err := db.GetDb()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could find games")
	}

	games := database.GetAllGames()
	gameSummaries := make([]model.GameSummary, 0)

	for _, g := range *games {
		gameSummaries = append(gameSummaries, model.GameToSummary(g))
	}

	return c.JSON(http.StatusOK, gameSummaries)
}

func newGame(c echo.Context) error {
	game, gameErr := createNewGame()

	if gameErr != nil {
		return gameErr
	}

	return c.JSON(http.StatusOK, buildGameState(game, "0"))
}

/*
// func login(c echo.Context) error {
// 	username := c.Param("username")

// 	database, err := db.GetDb()

// 	if err != nil {
// 		return err
// 	}

// 	player, playerErr := database.CreatePlayer(username)

// 	if playerErr != nil {
// 		return playerErr
// 	}

// 	//token, err := newJWT(username, player.ID)

// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, &Response{true, buildGameState(game, "0")})
// }

func join(c echo.Context) error {
	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
	player, validPlayer, err := getPlayerFromHeader(authHeader)

	if err != nil {
		return err
	}

	if !validPlayer {
		return c.JSON(http.StatusUnauthorized, &Response{false, nil})
	}

	game, err := joinGame(c.Param("game"), player)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{true, buildGameState(game, "0")})
}

func startGame(c echo.Context) error {
	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
	player, validPlayer, err := getPlayerFromHeader(authHeader)

	if err != nil {
		return err
	}

	if !validPlayer {
		return c.JSON(http.StatusUnauthorized, &Response{false, nil})
	}

	dealCards(c.Param("game"), player)
	return update(c)
}

func getGameState(c echo.Context) error {
	playerID := getPlayerFromContext(c)

	game, _ := updateGame(c.Param("game"), nil)

	return c.JSON(http.StatusOK, &Response{true, buildGameState(game, playerID)})
}

func update(c echo.Context) error {
	playerID := getPlayerFromContext(c)
	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
	player, validPlayer, err := getPlayerFromHeader(authHeader)

	if err != nil {
		return err
	}

	if !validPlayer {
		return c.JSON(http.StatusUnauthorized, &Response{false, nil})
	}

	game, err := updateGame(c.Param("game"), player)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{true, buildGameState(game, playerID)})
}

func play(c echo.Context) error {
	// TODO Cards have a value, which can include skip, reverse, etc
	playerID := getPlayerFromContext(c)
	card := model.Card{c.Param("number"), c.Param("color")}

	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
	player, validPlayer, err := getPlayerFromHeader(authHeader)

	if err != nil {
		return err
	}

	if !validPlayer {
		return c.JSON(http.StatusUnauthorized, &Response{false, nil})
	}

	game, err := playCard(c.Param("game"), player, card)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{true, buildGameState(game, playerID)})
}

func draw(c echo.Context) error {
	playerID := getPlayerFromContext(c)
	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
	player, validPlayer, err := getPlayerFromHeader(authHeader)

	if err != nil {
		return err
	}

	if !validPlayer {
		return c.JSON(http.StatusUnauthorized, &Response{false, nil})
	}

	game, err := drawCard(c.Param("game"), player)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{true, buildGameState(game, playerID)})

}
*/

func buildGameState(game *model.Game, playerID string) map[string]interface{} {
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
