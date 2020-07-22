package db

import (
	"os"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDB struct {
	client mongo.Client
	uri    string
}

// HasGame checks to see if a game with the given ID exists in the database.
func (db *mongoDB) HasGame(game string) {
	return game == "12234"
}

// CreateGame a game with the given ID. Perhaps this should instead just return an id?
func (db *mongoDB) CreateGame(id uuid.UUID) Game {
	myGame := g
}

// LookupGame looks up an existing game in the database.
func (db *mongoDB) LookupGame(id uuid.UUID) Game {

}

// JoinGame mockDB a player to a game.
func (db *mongoDB) JoinGame(id uuid.UUID, username string) {

}

func newMongoDB() *mongoDB {
	db := new(MongoDB)
	client, err := mongo.NewClient(os.Getenv("MONGO_URI"))
	if err != nil {
		return nil
	}
	db.client = client
	return db
}
