package db

import (
	"github.com/google/uuid"
	"github.com/jak103/uno/model"
)

// UnoDB declares the database types for the applicaiton
type UnoDB interface {
	// Check if a game with the given ID exists in the database.
	HasGame(game uuid.UUID) bool
	// Creates a game with the given ID. Perhaps this should instead just return an id?
	CreateGame(id uuid.UUID) model.Game
	// Looks up an existing game in the database.
	LookupGame(id uuid.UUID) model.Game
	// Joins a player to a game.
	JoinGame(id uuid.UUID, username string)
}
