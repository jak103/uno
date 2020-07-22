package db

import (
	"github.com/google/uuid"
	"github.com/jak103/uno/model"
)

type firebaseDB struct {
	uri string
}

// HasGame checks to see if a game with the given ID exists in the database.
func (db *firebaseDB) HasGame(game string) bool {
	return game == "12234"
}

// CreateGame a game with the given ID. Perhaps this should instead just return an id?
func (db *firebaseDB) CreateGame(id uuid.UUID) model.Game {
	myGame := model.Game{ID: uuid.Nil, Password: "12234"}
	return myGame
}

// LookupGame looks up an existing game in the database.
func (db *firebaseDB) LookupGame(id uuid.UUID) model.Game {
	myGame := model.Game{ID: uuid.Nil, Password: "12234"}
	return myGame
}

// JoinGame mockDB a player to a game.
func (db *firebaseDB) JoinGame(id uuid.UUID, username string) {
	return
}

func newFirebaseDB() *firebaseDB {
	db := new(firebaseDB)
	return db
}
