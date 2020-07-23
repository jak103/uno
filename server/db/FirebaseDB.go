package db

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jak103/uno/model"
)

type firebaseDB struct {
	games         map[string]model.Game
	gamePasswords map[string]model.Game
	players       map[string]model.Player
}

// HasGame checks to see if a game with the given ID exists in the database.
func (db *firebaseDB) HasGameByPassword(password string) bool {
	_, ok := db.gamePasswords[password]
	return ok
}

// HasGameByID checks to see if a game with the given ID exists in the database.
func (db *firebaseDB) HasGameByID(id string) bool {
	_, ok := db.games[id]
	return ok
}

// CreateGame a game with the given ID. Perhaps this should instead just return an id?
func (db *firebaseDB) CreateGame() (*model.Game, error) {
	myGame := model.Game{ID: uuid.New().String(), Password: "12234"}
	db.games[myGame.ID] = myGame
	db.gamePasswords[myGame.Password] = myGame
	return &myGame, nil
}

// DeleteGame deletes a game
func (db *firebaseDB) DeleteGame(id string) error {
	return nil
}

// CreatePlayer creates the player in the database
func (db *firebaseDB) CreatePlayer(name string) (*model.Player, error) {
	player := model.Player{ID: uuid.New().String(), Name: name}
	db.players[player.ID] = player
	return &player, nil
}

// DeletePlayer deletes a player from the database
func (db *firebaseDB) DeletePlayer(id string) error {
	return nil
}

// LookupGameByID looks up an existing game in the database.
func (db *firebaseDB) LookupGameByID(id string) (*model.Game, error) {
	if game, ok := db.games[id]; ok {
		return &game, nil
	}
	return nil, errors.New("mockdb: game not found")
}

// LookupGameByPassword looks up an existing game in the database.
func (db *firebaseDB) LookupGameByPassword(password string) (*model.Game, error) {
	if game, ok := db.gamePasswords[password]; ok {
		return &game, nil
	}
	return nil, errors.New("mockdb: game not found")
}

// LookupPlayer checks to see if a player is in the database
func (db *firebaseDB) LookupPlayer(id string) (*model.Player, error) {
	if player, ok := db.players[id]; ok {
		return &player, nil
	}
	return nil, errors.New("mockdb: player not found")
}

// JoinGame join a player to a game.
func (db *firebaseDB) JoinGame(id string, username string) error {
	return nil
}

// SaveGame saves the game
func (db *firebaseDB) SaveGame(game model.Game) error {
	db.games[game.ID] = game
	db.gamePasswords[game.Password] = game
	return nil
}

// SavePlayer saves the player data
func (db *firebaseDB) SavePlayer(player model.Player) error {
	db.players[player.ID] = player
	return nil
}

// Disconnect disconnects from the remote database
func (db *firebaseDB) disconnect() {
	return
}

// Connect allows the user to connect to the database
func (db *firebaseDB) connect() {
	return
}

func init() {
	registerDB(&DB{
		name:        "FIREBASE",
		description: "Production Firebase connection",
		UnoDB:       new(firebaseDB),
	})
}
