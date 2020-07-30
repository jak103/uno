package main

import (
	"errors"

	"github.com/jak103/uno/db"
	"github.com/jak103/uno/model"
)

////////////////////////////////////////////////////////////
// Utility functions used in place of firebase
////////////////////////////////////////////////////////////
func randColor(i int) string {
	switch i {
	case 0:
		return "red"
	case 1:
		return "blue"
	case 2:
		return "green"
	case 3:
		return "yellow"
	}
	return ""
}

////////////////////////////////////////////////////////////
// Utility functions
////////////////////////////////////////////////////////////

// TODO: make sure this reflects on the front end
// func checkForWinner(game *model.Game) string {
// 	for k := range game.Players {
// 		if len(allCards[players[k]]) == 0 {
// 			return players[k]
// 		}
// 	}
// 	return ""
// }

////////////////////////////////////////////////////////////
// These are all of the functions for the game -> essentially public functions
////////////////////////////////////////////////////////////
func updateGame(game string, username string) (*model.Game, error) {
	database, err := db.GetDb()

	if err != nil {
		return nil, err
	}

	gameData, gameErr := database.LookupGameByID(game)

	if gameErr != nil {
		return nil, err
	}

	found := false
	for i := 0; i < len(gameData.Players); i++ {
		player := gameData.Players[i]
		if player.Name == username {
			found = true
			break
		}
	}

	if !found {
		return nil, errors.New("Player not in game, cannot start")
	}

	if gameData.Status != "Playing" {
		gameData.Status = "Playing"
	}

	err = database.SaveGame(*gameData)

	if err != nil {
		return nil, err
	}

	return gameData, nil
}

func createNewGame() (*model.Game, error) {
	database, err := db.GetDb()

	if err != nil {
		return nil, err
	}

	game, err := database.CreateGame()

	if err != nil {
		return nil, err
	}

	return game, nil
}

func joinGame(game string, username string) (*model.Game, error) {
	database, err := db.GetDb()
	if err != nil {
		return nil, err
	}

	player, playerErr := database.CreatePlayer(username)

	if playerErr != nil {
		return nil, err
	}

	gameData, gameErr := database.JoinGame(game, player.ID)

	if gameErr != nil {
		return nil, gameErr
	}

	return gameData, nil
}

func playCard(game string, username string, card model.Card) (*model.Game, error) {
	database, err := db.GetDb()

	if err != nil {
		return nil, err
	}

	gameData, gameErr := database.LookupGameByID(game)

	if gameErr != nil {
		return nil, err
	}

	// if gameData.CurrentPlayer == username {
	// 	cards := allCards[username]
	// 	if card.Color == currCard[0].Color || card.Value == currCard[0].Value {
	// 		// Valid card can be played
	// 		playerIndex = (playerIndex + 1) % len(players)
	// 		currPlayer = players[playerIndex]
	// 		currCard[0] = card

	// 		for index, item := range cards {
	// 			if item == currCard[0] {
	// 				allCards[username] = append(cards[:index], cards[index+1:]...)
	// 				break
	// 			}
	// 		}
	// 	}
	// 	return true
	// }
	return gameData, nil
}

// TODO: Keep track of current card that is top of the deck
func drawCard(game string, username string) (*model.Game, error) {
	database, err := db.GetDb()

	if err != nil {
		return nil, err
	}

	gameData, gameErr := database.LookupGameByID(game)

	if gameErr != nil {
		return nil, err
	}

	// if checkID(game) && username == currPlayer {
	// 	playerIndex = (playerIndex + 1) % len(players)
	// 	currPlayer = players[playerIndex]
	// 	// TODO: Use deck utils instead
	// 	//allCards[username] = append(allCards[username], newRandomCard()[0])
	// 	return true
	// }
	return gameData, nil
}

// TODO: need to deal the actual cards, not just random numbers
func dealCards(game string, username string) (*model.Game, error) {
	database, err := db.GetDb()

	if err != nil {
		return nil, err
	}

	gameData, gameErr := database.LookupGameByID(game)

	if gameErr != nil {
		return nil, err
	}

	// The game has started, no more players are joining
	// loop through players, set their cards
	// gameStarted = true
	// currPlayer = players[rand.Intn(len(players))]
	// deck := generateShuffledDeck()

	// for k := range players {
	// 	cards := []model.Card{}
	// 	for i := 0; i < 7; i++ {

	// 		drawnCard := deck[len(deck)-1]
	// 		deck = deck[:len(deck)-1]
	// 		cards = append(cards, drawnCard)
	// 		//cards = append(cards, model.Card{rand.Intn(10), randColor(rand.Intn(4))})
	// 	}
	// 	allCards[players[k]] = cards
	// }

	// currCard = deck
	//currCard = newRandomCard()

	return gameData, nil
}
