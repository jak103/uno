package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jak103/uno/db"
	"github.com/jak103/uno/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var tokenSecret string = "usudevops"

func setupRoutes(e *echo.Echo) {
	// Routes that don't require a valid JWT
	e.GET("/api/games", getGames)
	e.POST("/api/games", newGame)
	e.POST("/api/games/:id/join", joinExistingGame)

	// Create a group that requires a valid JWT
	group := e.Group("/api")

	group.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(tokenSecret),
		AuthScheme: "Token",
	}))

	group.POST("/games/:id/start", startGame)
	group.POST("/games/:id/play", play) // Ryan Johnson
	group.POST("/games/:id/draw", draw) // Brady Svedin
	//	group.POST("/games/:id/uno", callUno)

	group.GET("/games/:id", getGameState)
}

func getGames(c echo.Context) error {
	//log.Println("Running getGames")
	database, err := db.GetDb()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not find games: Failed to connect to db")
	}

	games, err := database.GetAllGames()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not find games")
	}

	gameSummaries := make([]model.GameSummary, 0)
	for _, g := range *games {
		summary := model.GameToSummary(g)
		gameSummaries = append(gameSummaries, summary)
	}

	return c.JSON(http.StatusOK, gameSummaries)
}

func newGame(c echo.Context) error {
	m := echo.Map{}

	err := c.Bind(&m)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not bind to input")
	}

	if m["name"] == nil || m["creator"] == nil {
		return c.JSON(http.StatusBadRequest, "Missing game name or creator")
	}

	gameName := m["name"].(string)
	creatorName := m["creator"].(string)

	if gameName == "" || creatorName == "" {
		return c.JSON(http.StatusBadRequest, "Missing game name or creator")
	}

	game, creator, gameErr := createNewGame(gameName, creatorName)

	if gameErr != nil {
		return gameErr
	}

	// Create token
	token := generateToken(creator)

	return c.JSON(http.StatusOK, map[string]interface{}{"token": token, "game": buildGameState(game, creator.ID)})
}

func joinExistingGame(c echo.Context) error {
	gameID := c.Param("id")
	m := echo.Map{}
	err := c.Bind(&m)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could bind to input")
	}

	if m["playerName"] == nil {
		return c.JSON(http.StatusBadRequest, "Missing player name")
	}

	playerName := m["playerName"].(string)

	if playerName == "" {
		return c.JSON(http.StatusBadRequest, "Missing player name")
	}

	player, _ := createPlayer(playerName)

	game, _ := joinGame(gameID, player)

	token := generateToken(player)

	return c.JSON(http.StatusOK, map[string]interface{}{"token": token, "game": buildGameState(game, player.Name)})
}

func generateToken(p *model.Player) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["playerName"] = p.Name
	claims["playerId"] = p.ID
	claims["exp"] = time.Now().Add(time.Hour * 4).Unix()

	t, err := token.SignedString([]byte(tokenSecret))

	if err != nil {
		return ""
	}

	return t
}

func getGameState(c echo.Context) error {
	playerID := getPlayerFromContext(c)
	gameID := c.Param("id")
	log.Println("playerID", playerID)
	log.Println("gameID", gameID)

	game, err := getGameUpdate(gameID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid game ID")
	}

	return c.JSON(http.StatusOK, buildGameState(game, playerID))

	//return c.JSON(http.StatusOK, map[string]interface{}{"game": buildGameState(game, playerID)}) //buildGameState(game, playerID))
}

func startGame(c echo.Context) error {
	playerID := getPlayerFromContext(c)

	database, err := db.GetDb()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not connect to database.")
	}

	gameID := c.Param("id")

	game, gameErr := database.LookupGameByID(gameID)

	if gameErr != nil {
		return c.JSON(http.StatusInternalServerError, "Could not find game.")
	}

	if game.Creator.ID != playerID {
		return c.JSON(http.StatusUnauthorized, "Only the player who created the game can start it.")
	}

	// get the game state back after dealing cards, etc.
	game, saveErr := dealCards(game)

	if saveErr != nil {
		return c.JSON(http.StatusInternalServerError, "Could not save game state.")
	}

	gameState := buildGameState(game, playerID)

	return c.JSON(http.StatusOK, gameState)
}

func play(c echo.Context) error {
	playerID := getPlayerFromContext(c)
	var playCard model.PlayCard
	c.Bind(&playCard)

	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
	player, validPlayer, err := getPlayerFromHeader(authHeader)

	if !validPlayer {
		return c.JSON(http.StatusBadRequest, "Not a valid player!")
	}

	game, err := playCard(c.Param("game"), player, card)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error playing the game card")
	}

	return c.JSON(http.StatusOK, buildGameState(game, playerID))
}

func draw(c echo.Context) error {
	playerID := getPlayerFromContext(c)
	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
	player, validPlayer, err := getPlayerFromHeader(authHeader)

	if err != nil {
		return err
	}

	if !validPlayer {
		return c.JSON(http.StatusUnauthorized, "")
	}

	game, err := drawCard(c.Param("game"), player)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, buildGameState(game, playerID))

}

func buildGameState(game *model.Game, playerID string) map[string]interface{} {
	gameState := make(map[string]interface{})

	// Update known variables
	gameState["direction"] = game.Direction
	gameState["draw_pile"] = game.DrawPile
	gameState["discard_pile"] = game.DiscardPile
	gameState["game_id"] = game.ID
	gameState["status"] = game.Status
	gameState["name"] = game.Name
	gameState["player_id"] = playerID
	if game.DiscardPile != nil {
		gameState["current_card"] = game.DiscardPile[0]
	} else {
		gameState["current_card"] = model.Card{}
	}

	for _, player := range game.Players {
		if player.ID != playerID {
			for _, card := range player.Cards {
				card.Color = "Blank"
				card.Value = "Blank"
			}
		} else {
			gameState["player_cards"] = player.Cards
		}
	}

	gameState["all_players"] = game.Players
	gameState["current_player"] = game.Players[game.CurrentPlayer]

	return gameState
}

func getPlayerFromContext(c echo.Context) string {
	// TODO Update this to the actual claim key once the JWT team is done
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	playerID := claims["playerId"].(string)

	return playerID
}
