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

func (db *firebaseDB) CreatePlayer(name string) (*model.Player, error) {
	player := model.Player{ID: uuid.New().String(), Name: name}
	db.players[player.ID] = player
	return &player, nil
}

// LookupGame looks up an existing game in the database.
func (db *firebaseDB) LookupGameByID(id string) (*model.Game, error) {
	if game, ok := db.games[id]; ok {
		return &game, nil
	}
	return nil, errors.New("mockdb: game not found")
}

// LookupGame looks up an existing game in the database.
func (db *firebaseDB) LookupGameByPassword(password string) (*model.Game, error) {
	if game, ok := db.gamePasswords[password]; ok {
		return &game, nil
	}
	return nil, errors.New("mockdb: game not found")
}

func (db *firebaseDB) LookupPlayer(id string) (*model.Player, error) {
	if player, ok := db.players[id]; ok {
		return &player, nil
	}
	return nil, errors.New("mockdb: player not found")
}

// JoinGame firebaseDB a player to a game.
func (db *firebaseDB) JoinGame(id string, username string) error {
	return nil
}

func (db *firebaseDB) SaveGame(game model.Game) error {
	db.games[game.ID] = game
	db.gamePasswords[game.Password] = game
	return nil
}

func (db *firebaseDB) SavePlayer(player model.Player) error {
	db.players[player.ID] = player
	return nil
}

func newFirebaseDB() *firebaseDB {
	return new(firebaseDB)
}
