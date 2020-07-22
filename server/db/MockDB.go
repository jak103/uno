package db

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jak103/uno/model"
)

// MockDB is an implemenation declaring the unit test db
type mockDB struct {
	games         map[string]model.Game
	gamePasswords map[string]model.Game
	players       map[string]model.Player
}

//type mongoDB struct {
//	client   *mongo.Client
//	uri      string
//	database *mongo.Database
//	games    *mongo.Collection
//	players  *mongo.Collection
//}

// HasGame checks to see if a game with the given ID exists in the database.
func (db *mockDB) HasGameByPassword(password string) bool {
	_, ok := db.gamePasswords[password]
	return ok
}

func (db *mockDB) HasGameByID(id string) bool {
	_, ok := db.games[id]
	return ok
}

// CreateGame a game with the given ID. Perhaps this should instead just return an id?
func (db *mockDB) CreateGame() (*model.Game, error) {
	myGame := model.Game{ID: uuid.New().String(), Password: "12234"}
	db.games[myGame.ID] = myGame
	db.gamePasswords[myGame.Password] = myGame
	return &myGame, nil
}

func (db *mockDB) CreatePlayer(name string) (*model.Player, error) {
	player := model.Player{ID: uuid.New().String(), Name: name}
	db.players[player.ID] = player
	return &player, nil
}

// LookupGame looks up an existing game in the database.
func (db *mockDB) LookupGameByID(id string) (*model.Game, error) {
	if game, ok := db.games[id]; ok {
		return &game, nil
	}
	return nil, errors.New("mockdb: game not found")
}

// LookupGame looks up an existing game in the database.
func (db *mockDB) LookupGameByPassword(password string) (*model.Game, error) {
	if game, ok := db.gamePasswords[password]; ok {
		return &game, nil
	}
	return nil, errors.New("mockdb: game not found")
}

func (db *mockDB) LookupPlayer(id string) (*model.Player, error) {
	if player, ok := db.players[id]; ok {
		return &player, nil
	}
	return nil, errors.New("mockdb: player not found")
}

// JoinGame mockDB a player to a game.
func (db *mockDB) JoinGame(id string, username string) error {
	return nil
}

func (db *mockDB) SaveGame(game model.Game) error {
	db.games[game.ID] = game
	db.gamePasswords[game.Password] = game
	return nil
}

func (db *mockDB) SavePlayer(player model.Player) error {
	db.players[player.ID] = player
	return nil
}

func newMockDB() *mockDB {
	return new(mockDB)
}
