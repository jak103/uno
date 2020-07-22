package db

import (
	"github.com/jak103/uno/model"
)

// UnoDB declares the database types for the applicaiton
type UnoDB interface {
	// Check if a game with the given password exists in the database.
	HasGameByPassword(password string) bool
	// Check if a game with the given ID exists in the database.
	HasGameByID(game string) bool
	// Creates a game with the given ID. Perhaps this should instead just return an id?
	CreateGame() (*model.Game, error)
	// Looks up an existing game in the database.
	LookupGameByID(id string) (*model.Game, error)
	// Looks up an existing game in the database.
	LookupGameByPassword(password string) (*model.Game, error)
	// Joins a player to a game.
	JoinGame(id string, username string)
}
