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

func (db *mockDB) GetAllGames() (*[]model.Game, error) {
	games := make([]model.Game, 0)

	for _, game := range db.games {
		games = append(games, game)
	}

	return &games, nil
}

// HasGame checks to see if a game with the given ID exists in the database.
func (db *mockDB) HasGameByPassword(password string) bool {
	_, ok := db.gamePasswords[password]
	return ok
}

// HasGameByID checks to see if a game with the given ID exists in the database.
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

// CreatePlayer creates the player in the database
func (db *mockDB) CreatePlayer(name string) (*model.Player, error) {
	player := model.Player{ID: uuid.New().String(), Name: name}
	db.players[player.ID] = player
	return &player, nil
}

// DeleteGame deletes a game
func (db *mockDB) DeleteGame(id string) error {
	if _, ok := db.games[id]; ok {
		delete(db.games, id)
	}
	return nil
}

// DeletePlayer deletes a player from the database
func (db *mockDB) DeletePlayer(id string) error {
	if _, ok := db.players[id]; ok {
		delete(db.players, id)
	}
	return nil
}

// LookupGameByID looks up an existing game in the database.
func (db *mockDB) LookupGameByID(id string) (*model.Game, error) {
	if id != "" {
		if game, ok := db.games[id]; ok {
			return &game, nil
		}
	}

	return nil, errors.New("mockdb: game not found")
}

// LookupGameByPassword looks up an existing game in the database.
func (db *mockDB) LookupGameByPassword(password string) (*model.Game, error) {
	if game, ok := db.gamePasswords[password]; ok {
		return &game, nil
	}
	return nil, errors.New("mockdb: game not found")
}

// LookupPlayer checks to see if a player is in the database
func (db *mockDB) LookupPlayer(id string) (*model.Player, error) {
	if player, ok := db.players[id]; ok {
		return &player, nil
	}
	return nil, errors.New("mockdb: player not found")
}

// JoinGame join a player to a game.
func (db *mockDB) JoinGame(id string, username string) (*model.Game, error) {
	if game, ok := db.games[id]; ok {
		if player, err := db.LookupPlayer(username); err != nil {
			return nil, err
		} else {
			game.Players = append(game.Players, *player)
			return &game, nil
		}
	} else {
		return nil, errors.New("mockdb: game not found")
	}
}

// SaveGame saves the game
func (db *mockDB) SaveGame(game model.Game) error {
	db.games[game.ID] = game
	db.gamePasswords[game.Password] = game
	return nil
}

// SavePlayer saves the player data
func (db *mockDB) SavePlayer(player model.Player) error {
	db.players[player.ID] = player
	return nil
}

// Disconnect disconnects from the remote database
func (db *mockDB) disconnect() {
	return
}

// connect allows the user to connect to the database
func (db *mockDB) connect() {
	return
}

func init() {
	registerDB(&DB{
		name:        "MOCK",
		description: "Mock database connection for Unit Tests",
		UnoDB: &mockDB{
			games:         make(map[string]model.Game),
			gamePasswords: make(map[string]model.Game),
			players:       make(map[string]model.Player),
		},
	})
}
