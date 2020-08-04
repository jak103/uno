package main

import (
	"testing"
	//"github.com/jak103/uno/db"
	"github.com/jak103/uno/model"
	"github.com/stretchr/testify/assert"
)

// This function is meant to get a game and a player into the data base in a usable state for testing.
/*
func setupGameWithPlayer(database *db.DB) (*model.Game, *model.Player) {
	player, _ := database.CreatePlayer("Player 1")

	game, _ := database.CreateGame("Game 1", player.ID)

	game, _ = database.JoinGame(game.ID, player.ID)

	game.DrawPile = generateShuffledDeck(1)

	database.SaveGame(*game)

	return game, player
	// This is created "Game 1" games in the real db
	// TODO: This should use a MockDB
}
*/

func TestDrawCard(t *testing.T) {
	/*
	// Test passing in a bogus game id, we should get an error
	game, err := drawCard("Bogus game id", "Bogus player id")

	// Assert that we got an actual err
	assert.NotNil(t, err, "We did not error on a bogus game id")

	// Generate real game in database and real player
	database, _ := db.GetDb()
	game, player := setupGameWithPlayer(database)

	// Test Drawing a card with a full deck and real player
	game, err = drawCard(game.ID, player.ID)
	game, _ = database.LookupGameByID(game.ID)
	player = &game.Players[game.CurrentPlayer]

	// Assert that no error occured, the player has a new card and the draw pile
	// has one less card
	assert.Nil(t, err, "Failed to draw card.")
	assert.Equal(t, 1, len(player.Cards))
	assert.Equal(t, 107, len(game.DrawPile))

	// Move all cards into the discard pile, Empty out the draw pile completely,
	// and test drawing a card. It should resuffle leaving one card on the discard pile
	game.DiscardPile = append(game.DiscardPile, game.DrawPile...)
	game.DrawPile = game.DrawPile[:0]
	lastCard := game.DiscardPile[len(game.DiscardPile)-1]

	database.SaveGame(*game)

	game, err = drawCard(game.ID, player.ID)
	player = &game.Players[game.CurrentPlayer]

	//Assert no error, player has 2 cards from both draw tests,
	// draw is missing three from two draws and one in discard
	// discard has one last remaining card.
	// Assert last card in discard is actually to proper last card
	assert.Nil(t, err, "Failed to draw card.")
	assert.Equal(t, 2, len(player.Cards))
	assert.Equal(t, 105, len(game.DrawPile))
	assert.Equal(t, 1, len(game.DiscardPile))
	assert.Equal(t, lastCard.Color, game.DiscardPile[0].Color)
	assert.Equal(t, lastCard.Value, game.DiscardPile[0].Value)

	// Empty out both discard and draw piles, it should for now put a new deck on the draw pile.
	game.DrawPile = game.DrawPile[:0]
	game.DiscardPile = game.DiscardPile[:1]
	lastCard = game.DiscardPile[len(game.DiscardPile)-1]

	database.SaveGame(*game)

	game, err = drawCard(game.ID, player.ID)
	player = &game.Players[game.CurrentPlayer]

	// Assert no errors, assert player now has 3 cards
	// assert new draw pile with one missing
	// assert discard still has one card
	// Assert last card in discard is actually to proper last card
	assert.Nil(t, err, "Failed to draw card.")
	assert.Equal(t, 3, len(player.Cards))
	assert.Equal(t, 107, len(game.DrawPile))
	assert.Equal(t, 1, len(game.DiscardPile))
	assert.Equal(t, lastCard.Color, game.DiscardPile[0].Color)
	assert.Equal(t, lastCard.Value, game.DiscardPile[0].Value)

	// Create a bogus player with a bogus ID
	otherPlayer := model.Player{ID: " id 2 ", Name: "Name 2", Cards: []model.Card{}}

	// Simulate a someone trying to participate in a game they are not a part of.
	_, err = drawCard(game.ID, otherPlayer.ID)

	// Assert that we got an error from the draw card function as we should have.
	// Assert that the player didn't get any cards
	// Assert that the draw pile didn't lose any cards.
	assert.NotNil(t, err, "Player not in the game drew a card. Please make sure only players in the game can draw")
	assert.Equal(t, "You cannot participate in a game you do not belong", err.Error())
	assert.Equal(t, 0, len(otherPlayer.Cards))
	assert.Equal(t, 107, len(game.DrawPile))

	// Create a real player and add them to the game so there is more than one player.
	player2, _ := database.CreatePlayer("Player 2")

	game, _ = database.JoinGame(game.ID, player2.ID)

	database.SaveGame(*game)

	//Simulate drawing out of turn
	_, err = drawCard(game.ID, player2.ID)

	// Assert that we got an error from the draw card function as we should have.
	// Assert that the player didn't get any cards
	// Assert that the draw pile didn't lose any cards.
	assert.NotNil(t, err, "Player drew out of turn. Please make sure only the player whoes turn it is can play.")
	assert.Equal(t, "It is not your turn to play", err.Error())
	assert.Equal(t, 0, len(player2.Cards))
	assert.Equal(t, 107, len(game.DrawPile))
	*/
	// TODO: Once we actually support a MockDB for unit tests, check the validity of our response

}

func TestDealCards(t *testing.T) {
	/*
	// Generate real game in database and real player
	database, err := db.GetDb()
	game, player := setupGameWithPlayer(database)

	// Test Drawing a card with a full deck and real player
	game, err = dealCards(game)
	player = &game.Players[game.CurrentPlayer] //getting from the game who the current player is

	// Assert that no error occured, the player has a new card and the draw pile
	// has one less card
	assert.Nil(t, err, "Failed to deal cards.")
	assert.Equal(t, 7, len(player.Cards))
	assert.Equal(t, 100, len(game.DrawPile))
	assert.Equal(t, 1, len(game.DiscardPile))

	// Create additional players and add them to the game
	player2, _ := database.CreatePlayer("Player 2")
	player3, _ := database.CreatePlayer("Player 3")
	player4, _ := database.CreatePlayer("Player 4")
	player5, _ := database.CreatePlayer("Player 5")

	game, _ = database.JoinGame(game.ID, player2.ID)
	//Have to save in between each player being added or the game state wont recall any but the last
	database.SaveGame(*game)
	game, _ = database.JoinGame(game.ID, player3.ID)
	database.SaveGame(*game)
	game, _ = database.JoinGame(game.ID, player4.ID)
	database.SaveGame(*game)
	game, _ = database.JoinGame(game.ID, player5.ID)
	database.SaveGame(*game)

	//refresh the drawPile and the discardPile
	game.DrawPile = []model.Card{}
	game.DiscardPile = []model.Card{}

	// Test Drawing a card with a full deck and multiple players
	game, err = dealCards(game)
	// Assert that no error occured, the player has a new card and the draw pile
	// has one less card
	assert.Nil(t, err, "Failed to deal multiple players cards.")
	for _, player := range game.Players {
		assert.Equal(t, 7, len(player.Cards))
	}
	assert.Equal(t, 180, len(game.DrawPile))
	assert.Equal(t, 1, len(game.DiscardPile))
	*/
	// TODO: This test should be rewritten to use a MockDB
	assert.True(t, true)

}

func TestCheckForCardInHand(t *testing.T){
	//Created two cards One will be in the hand and the other won't
	validCard := model.Card{"red", "1"}
	falseCard := model.Card{"blue", "4"}
	//Created a hand with the valid card in it
	hand := []model.Card{validCard}
	
	//Testing to see if the function returns True for a card that is 
	//present and False for a card that isn't present
	assert.True(t, checkForCardInHand(validCard, hand))
	assert.False(t, checkForCardInHand(falseCard, hand))
}

func TestCreatePlayer(t *testing.T){
	/*
	// get the database
	database, err := db.GetDb()
	assert.Nil(t, err, "could not find database")
	// use the createPlayer function
	player, err := createPlayer("test")
	assert.Nil(t, err, "could not create player")
	// Lookup the player in the database to see if it is there
	databasePlayer, err := database.LookupPlayer(player.ID)
	assert.Nil(t, err, "could not find player")
	// Test to see if the database player and the created player are the same
	assert.Equal(t, player, databasePlayer)
	*/
	// TODO: This test should be rewritten to use a <ockDB
	assert.True(t, true)
}

func TestJoinGame(t *testing.T){
	/*
	// get database
	database, err := db.GetDb()
	assert.Nil(t, err, "could not find database")
	// create a new game with one player
	game, _, err := createNewGame("testGame", "testPlayer")
	assert.Nil(t, err, "could not create game")
	// create a new player
	newPlayer, err := createPlayer("joinGamePlayer")
	assert.Nil(t, err, "could not create new player")
	// attempt to join game
	game, err = joinGame(game.ID, newPlayer)
	database.SaveGame(*game)
	assert.Nil(t, err, "could not join game with new player")
	// lookup game from database 
	game, err = database.LookupGameByID(game.ID)
	assert.Nil(t, err, "could not find game in database")
	// test to see if the new Player is in the game
	assert.Contains(t, game.Players, *newPlayer)
	*/
	// TODO: This test should be rewritten to use a MockDB
	assert.True(t, true)
}


func TestDrawTopCard(t *testing.T) {
	/*
	game, _, err := createNewGame("Test Game", "Test Player")
	game.DrawPile = []model.Card{model.Card{"red", "1"}, model.Card{"blue", "2"}, model.Card{"green", "3"}}
	game, cardReturned := drawTopCard(game)
	assert.Equal(t, model.Card{"green", "3"}, cardReturned)
	*/
	// TODO: This test should be rewritten to use a MockDB
	assert.True(t, true)
}
