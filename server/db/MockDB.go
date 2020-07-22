package db

import (
	"github.com/google/uuid"
)

// MockDB is an implemenation declaring the unit test db
type mockDB struct{}

// HasGame checks to see if a game with the given ID exists in the database.
func (db *mockDB) HasGame(game string) bool {
	return game == "12234"
}

// CreateGame a game with the given ID. Perhaps this should instead just return an id?
func (db *mockDB) CreateGame(id uuid.UUID) Game {
	myGame := Game{ID: uuid.Nil, Password: "12234"}
	return myGame
}

// LookupGame looks up an existing game in the database.
func (db *mockDB) LookupGame(id uuid.UUID) Game {
	myGame := Game{ID: uuid.Nil, Password: "12234"}
	return myGame
}

// JoinGame mockDB a player to a game.
func (db *MockDB) JoinGame(id uuid.UUID, username string) {
	return
}

func newMockDB() *mockDB {
	return new(mockDB)
}
