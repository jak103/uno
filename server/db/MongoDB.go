package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jak103/uno/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	client   *mongo.Client
	uri      string
	database *mongo.Database
	games    *mongo.Collection
	players  *mongo.Collection
}

func (db *mongoDB) HasGameByID(id string) bool {
	return true
}

// HasGame checks to see if a game with the given ID exists in the database.
func (db *mongoDB) HasGameByPassword(password string) bool {
	game, err := db.LookupGameByPassword(password)
	return err == nil && game != nil
}

// CreateGame a game with the given ID. Perhaps this should instead just return an id?
func (db *mongoDB) CreateGame() (*model.Game, error) {
	myGame := model.Game{Password: ""}
	res, err := db.games.InsertOne(context.Background(), myGame)
	if err != nil {
		return nil, err
	}
	myGame.ID = res.InsertedID.(primitive.ObjectID).String()
	return &myGame, nil
}

// LookupGame looks up an existing game in the database.
func (db *mongoDB) LookupGameByPassword(password string) (*model.Game, error) {
	var game model.Game
	err := db.games.FindOne(context.Background(), bson.M{"password": password}).Decode(&game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

// LookupGame looks up an existing game in the database.
func (db *mongoDB) LookupGameByID(id string) (*model.Game, error) {
	var res model.Game
	if err := db.games.FindOne(context.Background(), bson.M{"_id": id}).Decode(&res); err != nil {
		return nil, err
	}
	fmt.Println(res)
	return &res, nil
}

// JoinGame mockDB a player to a game.
func (db *mongoDB) JoinGame(id string, username string) {
	return
}

func newMongoDB() *mongoDB {
	db := new(mongoDB)
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
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
	database := client.Database("uno")
	db.database = database
	db.games = database.Collection("games")
	db.players = database.Collection("players")
	return db
}
