package db

import (
	"context"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jak103/uno/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	client *mongo.Client
	uri    string
}

func (db *mongoDB) HasGameByID(id uuid.UUID) bool {
	return true
}

// HasGame checks to see if a game with the given ID exists in the database.
func (db *mongoDB) HasGameByPassword(password string) bool {
	return password == "12234"
}

// CreateGame a game with the given ID. Perhaps this should instead just return an id?
func (db *mongoDB) CreateGame(id uuid.UUID) model.Game {
	myGame := model.Game{ID: uuid.Nil, Password: "12234"}
	return myGame
}

// LookupGame looks up an existing game in the database.
func (db *mongoDB) LookupGameByPassword(password string) model.Game {
	myGame := model.Game{ID: uuid.Nil, Password: "12234"}
	return myGame
}

// LookupGame looks up an existing game in the database.
func (db *mongoDB) LookupGameByID(id uuid.UUID) model.Game {
	myGame := model.Game{ID: uuid.Nil, Password: "12234"}
	return myGame
}

// JoinGame mockDB a player to a game.
func (db *mongoDB) JoinGame(id uuid.UUID, username string) {
	return
}

func newMongoDB() *mongoDB {
	db := new(mongoDB)
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	db.client = client
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	return db
}
