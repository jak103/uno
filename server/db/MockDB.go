package db

import (
	"github.com/google/uuid"
	"github.com/jak103/uno/model"
)

// MockDB is an implemenation declaring the unit test db
type mockDB struct{
	games []model.Game
}

// HasGame checks to see if a game with the given ID exists in the database.
func (db *mockDB) HasGame(game string) bool {
	for i, myGame := range db.games {
		if myGame.ID == game {
			return true
		}
	}
	return false
}

// CreateGame a game with the given ID. Perhaps this should instead just return an id?
func (db *mockDB) CreateGame(id uuid.UUID) uuid.UUID {
	myGame := model.Game{ID: uuid.Nil, Password: "12234"}
	db = append(db.games, myGame)
	return myGame.ID
}

// LookupGame looks up an existing game in the database.
func (db *mockDB) LookupGame(id uuid.UUID) model.Game {
	myGame := model.Game{ID: uuid.Nil, Password: "12234"}
	return myGame
}

// JoinGame mockDB a player to a game.
func (db *mockDB) JoinGame(id uuid.UUID, username string) {
	return
}

func newMockDB() *mockDB {
	return new(mockDB)
}
