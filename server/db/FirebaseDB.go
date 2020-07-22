package db

import (
	"github.com/google/uuid"
	"github.com/jak103/uno/model"
)

type firebaseDB struct{}

// HasGame checks to see if a game with the given ID exists in the database.
func (db *firebaseDB) HasGameByPassword(password string) bool {
	return password == "12234"
}

func (db *firebaseDB) HasGameByID(id string) bool {
	return true
}

// CreateGame a game with the given ID. Perhaps this should instead just return an id?
func (db *firebaseDB) CreateGame() (*model.Game, error) {
	myGame := model.Game{ID: uuid.Nil.String(), Password: "12234"}
	return &myGame, nil
}

// LookupGame looks up an existing game in the database.
func (db *firebaseDB) LookupGameByID(id string) (*model.Game, error) {
	myGame := model.Game{ID: uuid.Nil.String(), Password: "12234"}
	return &myGame, nil
}

// LookupGame looks up an existing game in the database.
func (db *firebaseDB) LookupGameByPassword(password string) (*model.Game, error) {
	myGame := model.Game{ID: uuid.Nil.String(), Password: "12234"}
	return &myGame, nil
}

// JoinGame firebaseDB a player to a game.
func (db *firebaseDB) JoinGame(id string, username string) {
	return
}

func newFirebaseDB() *firebaseDB {
	return new(firebaseDB)
}
