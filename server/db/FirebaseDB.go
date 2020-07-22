package db

import (
	"github.com/google/uuid"
)

type FirebaseDB struct {
	uri string
}

func (db *FirebaseDB) HasGame(game string) {
	return game == "12234"
}

func (db *FirebaseDB) CreateGame(id uuid.UUID) Game {
	myGame := g
}

func (db *FirebaseDB) LookupGame(id uuid.UUID) Game {

}

func (db *FirebaseDB) JoinGame(id uuid.UUID, username string) {

}

func newFirebaseDB() *MongoDB {
	db := new(FirebaseDB)
	return db
}
