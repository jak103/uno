package main

import (
	"math/rand"

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
func getGameUpdate(gameID string) (*model.Game, error) {
	database, err := db.GetDb()

	if err != nil {
		return nil, err
	}

	gameData, gameErr := database.LookupGameByID(gameID)

	if gameErr != nil {
		return nil, err
	}

	return gameData, nil
}

func createPlayer(name string) (*model.Player, error) {
	database, err := db.GetDb()
	if err != nil {
		return nil, err
	}

	player, err := database.CreatePlayer(name)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func createNewGame(gameName string, creatorName string) (*model.Game, *model.Player, error) {
	database, err := db.GetDb()
	game, err := database.CreateGame()
	if err != nil {
		return nil, nil, err
	}

	creator, e := createPlayer(creatorName)

	if e != nil {
		return nil, nil, e
	}

	game, err = database.JoinGame(game.ID, creator.ID)
	if err != nil {
		return nil, nil, err
	}
	game.Name = gameName
	game.Creator = *creator
	game.Status = model.WaitingForPlayers

	err = database.SaveGame(*game)
	if err != nil {
		return nil, nil, err
	}

	return game, creator, nil
}

func joinGame(game string, player *model.Player) (*model.Game, error) {
	database, err := db.GetDb()
	if err != nil {
		return nil, err
	}

	gameData, gameErr := database.JoinGame(game, player.ID)

	if gameErr != nil {
		return nil, gameErr
	}

	return gameData, nil
}

func playCard(game string, player *model.Player, card model.Card) (*model.Game, error) {
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
func drawCard(game string, player *model.Player) (*model.Game, error) {
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
func dealCards(game *model.Game) (*model.Game, error) {

	// pick a starting player
	game.CurrentPlayer = rand.Intn(len(game.Players))

	// get a deck
	game.DrawPile = generateShuffledDeck()

	// give everyone a hand of seven cards
	for k := range game.Players {
		cards := []model.Card{}
		for i := 0; i < 7; i++ {

			drawnCard := game.DrawPile[len(game.DrawPile)-1]
			game.DrawPile = game.DrawPile[:len(game.DrawPile)-1]
			cards = append(cards, drawnCard)
		}
		game.Players[k].Cards = cards
	}

	// draw a card for the discard
	drawnCard := game.DrawPile[len(game.DrawPile)-1]
	game.DrawPile = game.DrawPile[:len(game.DrawPile)-1]

	game.DiscardPile = append(game.DiscardPile, drawnCard)

	game.Status = "Playing"

	database, err := db.GetDb()

	if err != nil {
		return nil, err
	}

	// save the new game status
	err = database.SaveGame(*game)

	return game, err
}
