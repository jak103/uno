package main

import (
	"errors"
	"fmt"

	"github.com/jak103/uno/db"
	"github.com/jak103/uno/model"
)

////////////////////////////////////////////////////////////
// Utility functions
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

// The function for playing a card. Right now it grabs the game, checks that the
// Player id exists in this game, then checks that they hold the card provided,
// If both are true it adds the card to the discard pile in the game and removes it
// From the players hand and we return true, else at the end we return false.
// We must do the checks because they are not done anywhere else.
func playCard(gameID string, playerID string, card model.Card) (*model.Game, error) {

	// These lines are simply getting the database and game and handling any error that could occur
	database, dbErr := db.GetDb()

	if dbErr != nil {
		return nil, dbErr
	}

	game, gameErr := database.LookupGameByID(gameID)

	if gameErr != nil {
		return nil, gameErr
	}

	//For loop that loops through all players in the game
	for _, player := range game.Players {
		// Check that currnt loop player has the matching id provided to function
		if playerID == player.ID {
			// Loop through all cards the player holds
			for index, item := range player.Cards {
				// check that they hold the card provided
				if item.Color == card.Color && item.Value == card.Value {
					//Remove the card from the players hand
					player.Cards = append(player.Cards[:index], player.Cards[index+1:]...)
					//add card to the discard pile
					game.DiscardPile = append(game.DiscardPile, card)
					// Save the game state
					database.SaveGame(*game)
					return game, nil
				}
			}
		}
	}
	// If you get here either the player did not exist in this game or
	// the player did not hold that card so we return false.

	// TODO Make real error
	return nil, fmt.Errorf("SOmething bad happend")

	// There are a couple ways this could be done.
	// We could use a helper function to get the player, instead of looking for it each time.
	/*
		player := getPlayer(game, playerId)

		if player == null{
			return false
		}

		for index, item := range player.Cards {
			if item.Color == card.Color && item.Value == card.Value {
				player.Cards = append(player.Cards[:index], player.Cards[index+1:]...)
				game.DiscardPile = append(game.DiscardPile, card)
				return true
			}
		}

	*/
}

func drawCard(gameID string, playerID string) (*model.Game, error) {
	// These lines are simply getting the database and game and handling any error that could occur
	database, dbErr := db.GetDb()

	if dbErr != nil {
		return nil, dbErr
	}

	gameData, gameErr := database.LookupGameByID(gameID)

	if gameErr != nil {
		return nil, gameErr
	}

	// We get the current player from the game
	player := &gameData.Players[gameData.CurrentPlayer]
	//We then check if the player attempting to play a card is the current player
	if player.ID == playerID {
		// We check if the draw pile has available cards
		if len(gameData.DrawPile) == 0 {
			// we check that the discard pile has cards to reshuffle
			if len(gameData.DiscardPile) <= 1 {
				// If there are not cards on the table add a new deck
				// TODO in the future do more complicated logic such as skip the players turn or something like that.
				gameData.DrawPile = generateShuffledDeck()
			} else {
				gameData = reshuffleDiscardPile(gameData)
			}
		}

		// Draw a card off the drawpile
		_, drawnCard := drawTopCard(gameData)

		// append the card into the players cards from the draw pile
		player.Cards = append(player.Cards, drawnCard)

		gameData = goToNextPlayer(gameData)

		// Save the game into the database
		database.SaveGame(*gameData)

		// Return a successfully updated game.
		return gameData, nil
	}

	// Check why they couldn't draw, is it not their turn, or are they not part of this game?
	for _, item := range gameData.Players {
		if item.ID == playerID {
			return nil, fmt.Errorf("It is not your turn to play")
		}
	}

	// TODO Make real error
	return nil, fmt.Errorf("You cannot participate in a game you do not belong")

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

func reshuffleDiscardPile(gameDate *model.Game) *model.Game {
	//Reshuffle all discarded cards except the last one back into the draw pile.
	oldDiscard := gameDate.DiscardPile[:len(gameDate.DiscardPile)-1]
	gameDate.DrawPile = shuffleCards(oldDiscard)
	gameDate.DiscardPile = gameDate.DiscardPile[len(gameDate.DiscardPile)-1:]
	return gameDate
}

func dealCards(gameID string, username string) (*model.Game, error) {
	database, err := db.GetDb()

	if err != nil {
		return nil, err
	}

	game, gameErr := database.LookupGameByID(gameID)

	if gameErr != nil {
		return nil, err
	}

	//TODO Match this up with the database.

	for k := range game.Players {
		cards := []model.Card{}
		for i := 0; i < 7; i++ {
			lastIndex := len(game.DrawPile) - 1
			card := game.DrawPile[lastIndex]
			cards = append(cards, card)
			game.DrawPile = game.DrawPile[:lastIndex]

		}
		game.Players[k].Cards = cards
	}
	//This will draw one more card, but instead of adding it to a players hand it will add it to the discard pile and set it as the current Card
	lastIndex := len(game.DrawPile) - 1
	startCard := game.DrawPile[lastIndex]
	game.DiscardPile = append(game.DiscardPile, startCard)
	game.DrawPile = game.DrawPile[:lastIndex]
	return game, nil
}

func drawTopCard(game *model.Game) (*model.Game, model.Card) {
	drawnCard := game.DrawPile[len(game.DrawPile)-1]
	game.DrawPile = game.DrawPile[:len(game.DrawPile)-1]
	return game, drawnCard
}
