package main

import (
	"errors"
	"fmt"

	"github.com/jak103/uno/db"
	"github.com/jak103/uno/model"
)

//Old Items wont need or use these anymore

////////////////////////////////////////////////////////////
// Utility functions
////////////////////////////////////////////////////////////

// A simple helper function to pull a card from a game and put it in the players hand.
// THis is used in  a lot of places, so this should be  a nice help
// Currently does not check for DrawPile size.
func drawCardHelper(game *model.Game, player *model.Player) {
	lastIndex := len(game.DrawPile) - 1
	card := game.DrawPile[lastIndex]

	player.Cards = append(player.Cards, card)
	game.DrawPile = game.DrawPile[:lastIndex]
}

// given a player and a card look for the card in players hand and return the index
// If it doesn't exists return -1
func cardFromPlayer(player *model.Player, card *model.Card) int {
	// Loop through all cards the player holds
	for index, item := range player.Cards {
		// check if current loop item matches card provided
		if item.Color == card.Color && item.Value == card.Value {
			// If the card matches return the current index
			return index
		}
	}
	// If we get to this point the player does not hold the card so we return nil
	return -1
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

	game, gameErr := database.LookupGameByID(gameID)

	if gameErr != nil {
		return nil, gameErr
	}

	// We get the current player from the game
	player := &game.Players[game.CurrentPlayer]
	//We then check if the player attempting to play a card is the current player
	if player.ID == playerID {
		// We check if the draw pile has available cards
		if len(game.DrawPile) == 0 {
			// we check that the discard pile has cards to reshuffle
			if len(game.DiscardPile) <= 1 {
				// If there are not cards on the table add a new deck
				// TODO in the future do more complicated logic such as skip the players turn or something like that.
				game.DrawPile = generateShuffledDeck()
			} else {
				//Reshuffle all discarded cards except the last one back into the draw pile.
				oldDiscard := game.DiscardPile[:len(game.DiscardPile)-1]
				game.DrawPile = shuffleCards(oldDiscard)
				game.DiscardPile = game.DiscardPile[len(game.DiscardPile)-1:]
			}
		}

		// Get the index of last card in draw pile, this is the card to be drawn.
		lastIndex := len(game.DrawPile) - 1

		// append the card into the players cards from the draw pile
		player.Cards = append(player.Cards, game.DrawPile[lastIndex])

		// Remove the card from the draw pile
		game.DrawPile = game.DrawPile[:lastIndex]

		// Save the game into the database
		database.SaveGame(*game)

		// Return a successfully updated game.
		return game, nil
	}

	// Check why they couldn't draw, is it not their turn, or are they not part of this game?
	for _, item := range game.Players {
		if item.ID == playerID {
			return nil, fmt.Errorf("It is not your turn to play")
		}
	}

	// TODO Make real error
	return nil, fmt.Errorf("You cannot participate in a game you do not belong")

}

/*This function will:
Deal out 7 cards to each player
Set the first card for the game to start from
*/
func dealCards(gameID string) (*model.Game, error) {
	database, err := db.GetDb()

	if err != nil {
		return nil, err
	}

	game, gameErr := database.LookupGameByID(gameID)

	if gameErr != nil {
		return nil, err
	}

	//For each player currently in the game
	for k := range game.Players {
		cards := []model.Card{}
		for i := 0; i < 7; i++ {

			//grab the top card from the draw pile
			lastIndex := len(game.DrawPile) - 1
			if lastIndex < 0 {
				//TODO check if something else should be happening, but for now we will just throw an extra deck on the pile
				game.DrawPile = append(game.DiscardPile, generateShuffledDeck()...)
				lastIndex = len(game.DrawPile) - 1
			}

			//append the card to the slice "cards" that we will add to the player
			cards = append(cards, game.DrawPile[lastIndex])
			//remove the top card from the draw pile
			game.DrawPile = game.DrawPile[:lastIndex]
		}
		//Add all 7 cards to that players hand
		game.Players[k].Cards = cards
	}
	//This will draw one more card, but instead of adding it to a players hand it will add it to the discard pile and set it as the current Card
	lastIndex := len(game.DrawPile) - 1
	game.DiscardPile = append(game.DiscardPile, game.DrawPile[lastIndex])
	game.DrawPile = game.DrawPile[:lastIndex]

	// Save the game into the database
	database.SaveGame(*game)

	return game, nil
}
