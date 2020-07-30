package db

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/jak103/uno/model"
	"google.golang.org/api/iterator"
)

type firestoreDB struct {
	client  *firestore.Client
	games   *firestore.CollectionRef
	players *firestore.CollectionRef
}

// HasGame checks to see if a game with the given ID exists in the database.
func (db *firestoreDB) HasGameByPassword(password string) bool {
	game, err := db.LookupGameByPassword(password)
	return err == nil && game != nil
}

// HasGameByID checks to see if a game with the given ID exists in the database.
func (db *firestoreDB) HasGameByID(id string) bool {
	game, err := db.LookupGameByID(id)
	return err == nil && game != nil
}

// CreateGame a game with the given ID. Perhaps this should instead just return an id?
func (db *firestoreDB) CreateGame() (*model.Game, error) {
	game := model.Game{ID: uuid.New().String(), Password: "12234"}
	gameDoc := db.games.Doc(game.ID)

	if _, err := gameDoc.Create(context.Background(), game); err != nil {
		return nil, err
	}

	return &game, nil
}

// CreatePlayer creates the player in the database
func (db *firestoreDB) CreatePlayer(name string) (*model.Player, error) {
	player := model.Player{ID: uuid.New().String(), Name: name}
	playerDoc := db.players.Doc(player.ID)

	if _, err := playerDoc.Create(context.Background(), player); err != nil {
		return nil, err
	}

	return &player, nil
}

// DeleteGame deletes a game
func (db *firestoreDB) DeleteGame(id string) error {
	gameDoc := db.games.Doc(id)

	if _, err := gameDoc.Delete(context.Background()); err != nil {
		return err
	}

	return nil
}

// DeletePlayer deletes a player from the database
func (db *firestoreDB) DeletePlayer(id string) error {
	playerDoc := db.players.Doc(id)

	if _, err := playerDoc.Delete(context.Background()); err != nil {
		return err
	}

	return nil
}

// LookupGameByID looks up an existing game in the database.
func (db *firestoreDB) LookupGameByID(id string) (*model.Game, error) {
	gameDoc := db.games.Doc(id)
	docSnapshot, err := gameDoc.Get(context.Background())

	if err != nil {
		return nil, err
	}

	if docSnapshot == nil {
		return nil, fmt.Errorf("%s: game not found", id)
	}

	var game model.Game
	if err = docSnapshot.DataTo(&game); err != nil {
		return nil, err
	}

	return &game, nil
}

// LookupGameByPassword looks up an existing game in the database.
func (db *firestoreDB) LookupGameByPassword(password string) (*model.Game, error) {
	q := db.games.Where("Password", "==", password)
	documents := q.Documents(context.Background())

	var game model.Game
	var err error
	var docSnapshot *firestore.DocumentSnapshot
	defer documents.Stop()
	for {
		docSnapshot, err = documents.Next()

		if err == iterator.Done {
			err = fmt.Errorf("%s: no game found for password", password)
			break
		}

		if err != nil {
			break
		}

		if err = docSnapshot.DataTo(&game); err != nil {
			break
		}

		break
	}

	if err != nil {
		return nil, err
	}

	return &game, nil
}

// LookupPlayer checks to see if a player is in the database
func (db *firestoreDB) LookupPlayer(id string) (*model.Player, error) {
	playerDoc := db.players.Doc(id)
	docSnapshot, err := playerDoc.Get(context.Background())

	if err != nil {
		return nil, err
	}

	if docSnapshot == nil {
		return nil, fmt.Errorf("%s: player not found", id)
	}

	var player model.Player
	if err = docSnapshot.DataTo(&player); err != nil {
		return nil, err
	}

	return &player, nil
}

// JoinGame join a player to a game.
func (db *firestoreDB) JoinGame(id string, username string) (*model.Game, error) {
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
func (db *firestoreDB) SaveGame(game model.Game) error {
	gameDoc := db.games.Doc(game.ID)

	if _, err := gameDoc.Set(context.Background(), game); err != nil {
		return err
	}

	return nil
}

// SavePlayer saves the player data
func (db *firestoreDB) SavePlayer(player model.Player) error {
	playerDoc := db.players.Doc(player.ID)

	if _, err := playerDoc.Set(context.Background(), player); err != nil {
		return err
	}

	return nil
}

// Disconnect disconnects from the remote database
func (db *firestoreDB) disconnect() {
	// Close the client connection if it is open
	if db.client != nil {
		defer db.client.Close()
	}
}

// Connect allows the user to connect to the database
func (db *firestoreDB) connect() {
	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("FIRESTORE_PROJECT_ID")

	client, err := firestore.NewClient(context.Background(), projectID)
	if err != nil {
		panic(err)
	}

	// Etablish Database Collection object
	db.client = client
	db.games = db.client.Collection("games")
	db.players = db.client.Collection("players")
}

func init() {
	registerDB(&DB{
		name:        "FIRESTORE",
		description: "Production Firestore connection",
		UnoDB:       new(firestoreDB),
	})
}
