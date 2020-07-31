package main

import (
	"errors"
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

func playCard(game string, playerID string, card model.Card) (*model.Game, error) {
	database, err := db.GetDb()

	if err != nil {
		return nil, err
	}

	gameData, gameErr := database.LookupGameByID(game)

	if gameErr != nil {
		return nil, err
	}

	if gameData.Players[gameData.CurrentPlayer].ID == playerID {
		hand := gameData.Players[gameData.CurrentPlayer].Cards
		if checkForCardInHand(card, hand) && (card.Color == gameData.DiscardPile[len(gameData.DiscardPile)-1].Color || card.Value == gameData.DiscardPile[len(gameData.DiscardPile)-1].Value || card.Value == "W4" || card.Value == "W") {
			// Valid card can be played

			gameData.DiscardPile = append(gameData.DiscardPile, card)

			for index, item := range hand {
				if item == card || (item.Value == "W" && card.Value == "W") || (item.Value == "W4" && card.Value == "W4") {
					gameData.Players[gameData.CurrentPlayer].Cards = append(hand[:index], hand[index+1:]...)
					break
				}
			}

			if card.Value == "S" {
				gameData = goToNextPlayer(gameData)
			}

			if card.Value == "D2" {
				gameData = goToNextPlayer(gameData)
				gameData = drawNCards(gameData, 2)
			}

			if card.Value == "W4" {
				gameData = goToNextPlayer(gameData)
				gameData = drawNCards(gameData, 4)
			}

			if card.Value == "R" {
				gameData.Direction = !gameData.Direction
			}

			gameData = goToNextPlayer(gameData)

		}
	}

	err = database.SaveGame(*gameData)

	if err != nil {
		return nil, err
	}

	return gameData, nil
}

func checkForCardInHand(card model.Card, hand []model.Card) bool {
	for _, c := range hand {
		// the wild cards, W4 and W, don't need to match in color; not for the previous card, and not with the hand. The card itself can become any color.
		if c.Value == card.Value && (c.Color == card.Color || card.Value == "W4" || card.Value == "W") {
			return true
		}
	}
	return false
}

// TODO: Keep track of current card that is top of the deck
func drawCard(gameID, playerID string) (*model.Game, error) {
	database, err := db.GetDb()

	if err != nil {
		return nil, err
	}

	gameData, gameErr := database.LookupGameByID(gameID)

	if gameErr != nil {
		return nil, err
	}

	if gameData.Players[gameData.CurrentPlayer].ID != playerID {
		return nil, errors.New("Wrong player")
	}

	var drawnCard model.Card
	gameData, drawnCard = drawTopCard(gameData)
	gameData.Players[gameData.CurrentPlayer].Cards = append(gameData.Players[gameData.CurrentPlayer].Cards, drawnCard)

	gameData = goToNextPlayer(gameData)

	database.SaveGame(*gameData)

	return gameData, nil
}

func goToNextPlayer(gameData *model.Game) *model.Game {
	if gameData.Direction {
		gameData.CurrentPlayer++
		gameData.CurrentPlayer %= len(gameData.Players)
	} else {
		gameData.CurrentPlayer--
		if gameData.CurrentPlayer < 0 {
			gameData.CurrentPlayer = len(gameData.Players) - 1
		}
	}

	return gameData
}

func drawNCards(gameData *model.Game, nCards uint) *model.Game {
	for i := uint(0); i < nCards; i++ {
		var drawnCard model.Card
		gameData, drawnCard = drawTopCard(gameData)
		gameData.Players[gameData.CurrentPlayer].Cards = append(gameData.Players[gameData.CurrentPlayer].Cards, drawnCard)
	}
	return gameData
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
			var drawnCard model.Card
			game, drawnCard = drawTopCard(game)
			cards = append(cards, drawnCard)
		}
		game.Players[k].Cards = cards
	}

	// draw a card for the discard
	var drawnCard model.Card
	game, drawnCard = drawTopCard(game)
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

func drawTopCard(game *model.Game) (*model.Game, model.Card) {
	drawnCard := game.DrawPile[len(game.DrawPile)-1]
	game.DrawPile = game.DrawPile[:len(game.DrawPile)-1]
	return game, drawnCard
}

func checkGameExists(gameID string) (bool, error) {
	database, err := db.GetDb()

	if err != nil {
		return false, err
	}

	_, gameErr := database.LookupGameByID(gameID)

	if gameErr != nil {
		return false, gameErr
	}

	return true, nil
}
