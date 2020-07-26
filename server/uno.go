package main

import (
	"math/rand"

	"github.com/jak103/uno/db"
	"github.com/jak103/uno/model"
)

//Old Items wont need or use these anymore

// ////////////////////////////////////////////////////////////
// // Utility functions used in place of firebase
// ////////////////////////////////////////////////////////////
// func randColor(i int) string {
// 	switch i {
// 	case 0:
// 		return "red"
// 	case 1:
// 		return "blue"
// 	case 2:
// 		return "green"
// 	case 3:
// 		return "yellow"
// 	}
// 	return ""
// }

// ////////////////////////////////////////////////////////////
// // All the data needed for a simulation of the game
// // eventually, this will be replaced with firebase
// ////////////////////////////////////////////////////////////
// var gameID string = ""
// var currCard []model.Card = nil // The cards are much easier to render as a list
// var players []string = []string{}
// var playerIndex = 0 // Used to iterate through the players
// var currPlayer string = ""
// var allCards map[string][]model.Card = make(map[string][]model.Card) // k: username, v: list of cards
// var gameStarted bool = false

// func newRandomCard() []model.Card {
// TODO use deck utils instead
// 	return []model.Card{model.Card{rand.Intn(10), randColor(rand.Intn(4))}}
// }

////////////////////////////////////////////////////////////
// Utility functions
////////////////////////////////////////////////////////////

// A simple helper function to pull a card from a game and put it in the players hand.
// THis is used in  a lot of places, so this should be  a nice help
func drawCardHelper(game *model.Game, player *model.Player) {
	lastIndex := len(game.DrawPile) - 1
	card := game.DrawPile[lastIndex]

	player.Cards = append(player.Cards, card)
	game.DrawPile = game.DrawPile[:lastIndex]
}

// A simpler helper function for getting the player with a matching ID to playerID
// from the list of players in the game.
func getPlayer(game *model.Game, playerID string) *model.Player {
	for _, item := range game.Players {
		if playerID == item.ID {
			return &item
		}
	}
	return nil
}

func newPayload(user string) map[string]interface{} { // User will default to "" if not passed
	payload := make(map[string]interface{})

	// Return the game model instead of these individually.

	// Update known variables
	// payload["current_card"] = currCard
	// payload["current_player"] = currPlayer
	// payload["all_players"] = players
	// payload["deck"] = allCards[user] // returns nil if currPlayer = "" or user not in allCards
	// payload["game_id"] = gameID
	// payload["game_over"] = checkForWinner()

	return payload
}

func checkID(id string) bool {
	return id == gameID
}

func contains(arr []string, val string) (int, bool) {
	for i, item := range arr {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

////////////////////////////////////////////////////////////
// These are all of the functions for the game -> essentially public functions
////////////////////////////////////////////////////////////
func updateGame(game string, username string) bool {
	success := false
	if success = checkID(game); success && gameStarted {
		return true
	}
	return false
}

func createNewGame() error {
	database, err := db.GetDb()

	if err != nil {
		return err
	}

	game, err := database.CreateGame()

	if err != nil {
		return err
	}

	gameID = game.ID
	return nil
}

func joinGame(game string, username string) bool {
	if checkID(game) {
		user := username

		if _, found := contains(players, user); !found {
			players = append(players, user)
			allCards[user] = nil // No cards yet
		}
		return true
	}
	return false // bad game_id
}

// The function for playing a card. Right now it grabs the game, checks that the
// Player id exists in this game, then checks that they hold the card provided,
// If both are true it adds the card to the discard pile in the game and removes it
// From the players hand and we return true, else at the end we return false.
// We must do the checks because they are not done anywhere else.
func playCard(gameID string, playerID string, card model.Card) bool {
	database, dbErr := db.GetDb()

	if dbErr != nil {
		return false
	}

	game, gameErr := database.LookupGameByID(gameID)

	if gameErr != nil {
		return false
	}

	for _, player := range game.Players {
		if playerID == player.ID {
			for index, item := range player.Cards {
				if item.Color == card.Color && item.Value == card.Value {
					player.Cards = append(player.Cards[:index], player.Cards[index+1:]...)
					game.DiscardPile = append(game.DiscardPile, card)
					return true
				}
			}
		}
	}

	return false

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

// TODO: Keep track of current card that is top of the deck
func drawCard(gameID string, playerID string) bool {
	database, err := db.GetDb()

	if err != nil {
		return false
	}

	game, err := database.LookupGameByID(gameID)
	var player *model.Player
	for _, item := range game.Players {
		if playerID == item.ID {
			player = &item
		}
	}
	if player == nil {
		return false
	}
	drawCardHelper(game, player)

	database.SaveGame(*game)

	return true
}

// TODO: need to deal the actual cards, not just random numbers
func dealCards() {
	// The game has started, no more players are joining
	// loop through players, set their cards
	gameStarted = true
	currPlayer = players[rand.Intn(len(players))]

	for k := range players {
		cards := []model.Card{}
		for i := 0; i < 7; i++ {

			// TODO Use deck utils instead
			//cards = append(cards, model.Card{rand.Intn(10), randColor(rand.Intn(4))})
		}
		allCards[players[k]] = cards
	}

	// TODO Use deck utils instead
	//currCard = newRandomCard()
}

// TODO: make sure this reflects on the front end
func checkForWinner() string {
	for k := range players {
		if len(allCards[players[k]]) == 0 {
			return players[k]
		}
	}
	return ""
}
