package db

import (
	"os"

	"github.com/google/uuid"
	"github.com/jak103/uno/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDB struct {
	client mongo.Client
	uri    string
}

// HasGame checks to see if a game with the given ID exists in the database.
func (db *mongoDB) HasGame(game string) bool {
	return game == "12234"
}

// CreateGame a game with the given ID. Perhaps this should instead just return an id?
func (db *mongoDB) CreateGame(id uuid.UUID) model.Game {
	myGame := model.Game{ID: uuid.Nil, Password: "12234"}
	return myGame
}

// LookupGame looks up an existing game in the database.
func (db *mongoDB) LookupGame(id uuid.UUID) model.Game {
	myGame := model.Game{ID: uuid.Nil, Password: "12234"}
	return myGame
}

// JoinGame mockDB a player to a game.
func (db *mongoDB) JoinGame(id uuid.UUID, username string) {
	return
}

func newMongoDB() *mongoDB {
	db := new(mongoDB)
	client, err := mongo.NewClient(os.Getenv("MONGO_URI"))
	if err != nil {
		return nil
	}
	db.client = client
	return db
}
