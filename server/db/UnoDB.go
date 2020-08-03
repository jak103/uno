package db

import (
	"github.com/jak103/uno/model"
)

// UnoDB declares the database types for the applicaiton
type UnoDB interface {
	// Returns all games in the database
	GetAllGames() (*[]model.Game, error)
	// Check if a game with the given password exists in the database.
	HasGameByPassword(password string) bool
	// Check if a game with the given ID exists in the database.
	HasGameByID(game string) bool
	// Creates a game.
	CreateGame(gameName string, creatorID string) (*model.Game, error)
	// Creates a player with the given name.
	CreatePlayer(name string) (*model.Player, error)
	// DeleteGame deletes a game
	DeleteGame(id string) error
	// DeletePlayer deletes a player from the database
	DeletePlayer(id string) error
	// Looks up an existing game in the database.
	LookupGameByID(id string) (*model.Game, error)
	// Looks up an existing game in the database.
	LookupGameByPassword(password string) (*model.Game, error)
	// Looks up an existing player in the database.
	LookupPlayer(id string) (*model.Player, error)
	// Joins a player to a game.
	JoinGame(gameID string, playerID string) (*model.Game, error)
	// Saves a game to the database.
	SaveGame(model.Game) error
	// Saves a player to the database.
	SavePlayer(model.Player) error
	// Adds a Players message to the db
	AddMessage(gameID string, playerID string, message model.Message) (*model.Game, error)
	// disconnects from the database.
	disconnect()
	// connect to the database
	connect()
}
