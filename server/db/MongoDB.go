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

func (db *mongoDB) GetAllGames() (*[]model.Game, error) {
	games := make([]model.Game, 0)

	cursor, err := db.games.Find(context.Background(), bson.M{}, nil)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		g := model.Game{}
		err := cursor.Decode(&g)
		if err != nil {
			panic(err)
		}
		games = append(games, g)
	}

	return &games, nil
}

// HasGame checks to see if a game with the given ID exists in the database.
func (db *mongoDB) HasGameByPassword(password string) bool {
	game, err := db.LookupGameByPassword(password)
	return err == nil && game != nil
}

// HasGameByID checks to see if a game with the given ID exists in the database.
func (db *mongoDB) HasGameByID(id string) bool {
	game, err := db.LookupGameByID(id)
	return err == nil && game != nil
}

// CreateGame a game with the given ID. Perhaps this should instead just return an id?
func (db *mongoDB) CreateGame() (*model.Game, error) {
	myGame := model.Game{Password: ""}
	res, err := db.games.InsertOne(context.Background(), myGame)
	if err != nil {
		return nil, err
	}
	myGame.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return &myGame, nil
}

// CreatePlayer creates the player in the database
func (db *mongoDB) CreatePlayer(name string) (*model.Player, error) {
	player := model.Player{Name: name}
	res, err := db.players.InsertOne(context.Background(), player)
	if err != nil {
		return nil, err
	}
	player.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return &player, nil
}

// DeleteGame deletes a game
func (db *mongoDB) DeleteGame(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = db.games.DeleteOne(context.Background(), bson.M{"_id": oid})
	if err != nil {
		return err
	}

	return nil
}

// DeletePlayer deletes a player from the database
func (db *mongoDB) DeletePlayer(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = db.players.DeleteOne(context.Background(), bson.M{"_id": oid})
	if err != nil {
		return err
	}

	return nil
}

// LookupGameByID looks up an existing game in the database.
func (db *mongoDB) LookupGameByID(id string) (*model.Game, error) {
	var res model.Game
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := db.games.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// LookupGameByPassword looks up an existing game in the database.
func (db *mongoDB) LookupGameByPassword(password string) (*model.Game, error) {
	var game model.Game
	err := db.games.FindOne(context.Background(), bson.M{"password": password}).Decode(&game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

// LookupPlayer checks to see if a player is in the database
func (db *mongoDB) LookupPlayer(id string) (*model.Player, error) {
	var res model.Player
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := db.players.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// JoinGame join a player to a game.
func (db *mongoDB) JoinGame(id string, username string) (*model.Game, error) {
	game, gameErr := db.LookupGameByID(id)

	if gameErr != nil {
		return nil, gameErr
	}

	player, playerErr := db.LookupPlayer(username)

	if playerErr != nil {
		return nil, playerErr
	}

	game.Players = append(game.Players, *player)

	err := db.SaveGame(*game)

	if err != nil {
		return nil, err
	}

	return game, nil
}

// SaveGame saves the game
func (db *mongoDB) SaveGame(game model.Game) error {
	savedID := game.ID
	defer func() { game.ID = savedID }()
	id, _ := primitive.ObjectIDFromHex(game.ID)
	game.ID = "" // Prevent Mongo from trying to change the ID.
	_, err := db.games.ReplaceOne(
		context.Background(),
		bson.M{"_id": id},
		game)
	return err
}

// SavePlayer saves the player data
func (db *mongoDB) SavePlayer(player model.Player) error {
	savedID := player.ID
	defer func() { player.ID = savedID }()
	id, _ := primitive.ObjectIDFromHex(player.ID)
	player.ID = "" // Prevent mongo from trying to change the ID.
	_, err := db.players.ReplaceOne(
		context.Background(),
		bson.M{"_id": id},
		player)
	return err
}

// disconnect disconnects from the remote database
func (db *mongoDB) disconnect() {
	fmt.Println("Disconnecting from the database.")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.client.Disconnect(ctx); err != nil {
		panic(err)
	}
	defer cancel()
}

// connect allows the user to connect to the database
func (db *mongoDB) connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	db.client = client
	database := client.Database("uno")
	db.database = database
	db.games = database.Collection("games")
	db.players = database.Collection("players")
}

func init() {
	registerDB(&DB{
		name:        "MONGO",
		description: "Mongo database for dev connections",
		UnoDB:       new(mongoDB),
	})
}
