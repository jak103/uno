package main

import (
	"testing"

	"github.com/jak103/uno/db"
	"github.com/jak103/uno/model"
	"github.com/stretchr/testify/assert"
)

// This function is meant to get a game and a player into the data base in a usable state for testing.
func setupGameWithPlayer(database *db.DB) (*model.Game, *model.Player) {

	player, _ := database.CreatePlayer("Player 1")

	game, _ := database.CreateGame("Game 1", player.ID)

	game, _ = database.JoinGame(game.ID, player.ID)

	game.DrawPile = generateShuffledDeck(1)

	game.CurrentPlayer = 0

	database.SaveGame(*game)

	return game, player
}

// Set up a game with 4 players, all with 7 known cards
// a full deck
// Direction is true
// top of discard is red 0
// current is 0
func setUpTestPlayerCard() *model.Game {

	database, _ := db.GetDb()

	player1, _ := database.CreatePlayer("Player 1")

	//Set up real game with 4 players and 7 known cards in each hand
	game, _ := database.CreateGame("Game 1", player1.ID)
	game, _ = database.JoinGame(game.ID, player1.ID)
	database.SaveGame(*game)

	player2, _ := database.CreatePlayer("Player 2")
	game, _ = database.JoinGame(game.ID, player2.ID)
	database.SaveGame(*game)

	player3, _ := database.CreatePlayer("Player 3")
	game, _ = database.JoinGame(game.ID, player3.ID)
	database.SaveGame(*game)

	player4, _ := database.CreatePlayer("Player 4")
	game, _ = database.JoinGame(game.ID, player4.ID)
	database.SaveGame(*game)

	game.DrawPile = generateShuffledDeck(1)
	game.CurrentPlayer = 0
	game.DiscardPile = []model.Card{model.Card{Color: "red", Value: "0"}}
	game.Direction = true

	game.Players[0].Cards = []model.Card{
		model.Card{Color: "red", Value: "2"},
		model.Card{Color: "red", Value: "R"},
		model.Card{Color: "red", Value: "D2"},
		model.Card{Color: "blue", Value: "6"},
		model.Card{Color: "blue", Value: "7"},
		model.Card{Color: "blue", Value: "9"},
		model.Card{Color: "yellow", Value: "8"},
	}

	game.Players[1].Cards = []model.Card{
		model.Card{Color: "red", Value: "S"},
		model.Card{Color: "red", Value: "3"},
		model.Card{Color: "green", Value: "6"},
		model.Card{Color: "green", Value: "4"},
		model.Card{Color: "blue", Value: "3"},
		model.Card{Color: "blue", Value: "0"},
		model.Card{Color: "blue", Value: "8"},
	}

	game.Players[2].Cards = []model.Card{
		model.Card{Color: "red", Value: "D2"},
		model.Card{Color: "red", Value: "2"},
		model.Card{Color: "blue", Value: "9"},
		model.Card{Color: "blue", Value: "5"},
		model.Card{Color: "Black", Value: "W4"},
		model.Card{Color: "blue", Value: "7"},
		model.Card{Color: "green", Value: "8"},
	}

	game.Players[3].Cards = []model.Card{
		model.Card{Color: "black", Value: "W"},
		model.Card{Color: "red", Value: "3"},
		model.Card{Color: "red", Value: "D2"},
		model.Card{Color: "green", Value: "5"},
		model.Card{Color: "blue", Value: "6"},
		model.Card{Color: "blue", Value: "7"},
		model.Card{Color: "yellow", Value: "8"},
	}

	database.SaveGame(*game)
	return game
}

func TestPlayCard_BadGame(t *testing.T) {
	// Starting by passing in bogus information to make sure we get proper errors.
	// Bogus game.
	game, err := playCard("Bogus Game ID", "Bogus Player ID", model.Card{Color: "Bogus Color", Value: "Bogus Value"})

	//Assert that the error is not nil, that it was returned
	// Assert we didn't get a real game back from the bogus id
	assert.NotNil(t, err)
	assert.Nil(t, game)
}
func TestPlayCard_BadPlayer(t *testing.T) {

	database, err := db.GetDb()

	game := setUpTestPlayerCard()

	gameID := game.ID

	//Real game, bogus player
	game, err = playCard(game.ID, "Bogus Player ID", model.Card{Color: "Bogus Color", Value: "Bogus Value"})

	// Assert there was an error, that the error is what we expect,
	// Assert we didn't get a game back.
	assert.NotNil(t, err)
	assert.Equal(t, "You cannot participate in a game you do not belong", err.Error())
	assert.Nil(t, game)

	//Pull game from database to do checks that it wasn't affected.
	game, _ = database.LookupGameByID(gameID)

	// Assert that the discard pile does not have a new card on it.
	// Assert that the two players still have all their cards.
	// Assert that the player turn did not change
	assert.Equal(t, 1, len(game.DiscardPile))
	assert.Equal(t, 7, len(game.Players[0].Cards))
	assert.Equal(t, 7, len(game.Players[1].Cards))
	assert.Equal(t, 7, len(game.Players[2].Cards))
	assert.Equal(t, 7, len(game.Players[3].Cards))
	assert.Equal(t, 0, game.CurrentPlayer)
}

func TestPlayCard_BogusCard(t *testing.T) {

	database, err := db.GetDb()

	game := setUpTestPlayerCard()

	gameID := game.ID

	// Test with real game, real player but bogus card not in players hand
	game, err = playCard(game.ID, game.Players[0].ID, model.Card{Color: "yellow", Value: "0"})

	// Assert error was returned
	// Assert proper error was returned
	// Assert game was not returned
	assert.NotNil(t, err)
	assert.Equal(t, "You cannot play a card that you do not own", err.Error())
	assert.Nil(t, game)

	//Pull game from database to do checks that it wasn't affected.
	game, _ = database.LookupGameByID(gameID)

	// Assert that the discard pile does not have a new card on it.
	// Assert that the two players still have all their cards.
	// Assert that the player turn did not change
	assert.Equal(t, 1, len(game.DiscardPile))
	assert.Equal(t, 7, len(game.Players[0].Cards))
	assert.Equal(t, 7, len(game.Players[1].Cards))
	assert.Equal(t, 7, len(game.Players[2].Cards))
	assert.Equal(t, 7, len(game.Players[3].Cards))
	assert.Equal(t, 0, game.CurrentPlayer)
}

func TestPlayCard_NotMatchingCard(t *testing.T) {

	database, err := db.GetDb()

	game := setUpTestPlayerCard()

	gameID := game.ID

	// Test with real game, real player but bogus card not in players hand
	game, err = playCard(game.ID, game.Players[0].ID, model.Card{Color: "blue", Value: "6"})

	// Assert error was returned
	// Assert proper error was returned
	// Assert game was not returned
	assert.NotNil(t, err)
	assert.Equal(t, "Your card must be either the same color or the same value as the top of the dicard pile or a wild", err.Error())
	assert.Nil(t, game)

	//Pull game from database to do checks that it wasn't affected.
	game, _ = database.LookupGameByID(gameID)

	// Assert that the discard pile does not have a new card on it.
	// Assert that the two players still have all their cards.
	// Assert that the player turn did not change
	assert.Equal(t, 1, len(game.DiscardPile))
	assert.Equal(t, 7, len(game.Players[0].Cards))
	assert.Equal(t, 7, len(game.Players[1].Cards))
	assert.Equal(t, 7, len(game.Players[2].Cards))
	assert.Equal(t, 7, len(game.Players[3].Cards))
	assert.Equal(t, 0, game.CurrentPlayer)
}

func TestPlayCard_NormalPlay(t *testing.T) {

	var err error
	game := setUpTestPlayerCard()

	//Set card to play that exists in player1 hand
	cardToPlay := model.Card{Color: "red", Value: "2"}
	oldCurrentPlayer := game.CurrentPlayer

	// Play with real game, real player, and real regular
	game, err = playCard(game.ID, game.Players[0].ID, cardToPlay)

	// Assert we did not get an error.
	// Assert discard pile has a new card
	// Assert top card is card played
	// Assert player1 has 1 less card
	// Assert current player was shifted up one
	// Assert proper card was removed from players hand
	assert.Nil(t, err)
	assert.Equal(t, 2, len(game.DiscardPile))
	assert.Equal(t, cardToPlay.Color, game.DiscardPile[len(game.DiscardPile)-1].Color)
	assert.Equal(t, cardToPlay.Value, game.DiscardPile[len(game.DiscardPile)-1].Value)
	assert.Equal(t, 6, len(game.Players[oldCurrentPlayer].Cards))
	assert.Equal(t, 1, game.CurrentPlayer)
	for _, card := range game.Players[oldCurrentPlayer].Cards {
		if card.Color == cardToPlay.Color && card.Value == cardToPlay.Value {
			assert.Fail(t, "Wrong card was removed from the players hand.")
		}
	}

}

func TestPlayCard_SkipCard(t *testing.T) {

	database, err := db.GetDb()

	game := setUpTestPlayerCard()

	//Set card to play that exists in player2 hand
	cardToPlay := model.Card{Color: "red", Value: "S"}
	game.CurrentPlayer = 1
	oldCurrentPlayer := 1

	database.SaveGame(*game)

	// Play a skip card
	game, err = playCard(game.ID, game.Players[1].ID, cardToPlay)

	// Assert we did not get an error.
	// Assert discard pile has a new card
	// Assert top card is card played
	// Assert player2 has 1 less card
	// Assert current player was shifted by two, so it should be player4 turn
	// Assert proper card was removed from players hand
	assert.Nil(t, err)
	assert.Equal(t, 2, len(game.DiscardPile))
	assert.Equal(t, cardToPlay.Color, game.DiscardPile[len(game.DiscardPile)-1].Color)
	assert.Equal(t, cardToPlay.Value, game.DiscardPile[len(game.DiscardPile)-1].Value)
	assert.Equal(t, 6, len(game.Players[oldCurrentPlayer].Cards))
	assert.Equal(t, 3, game.CurrentPlayer)
	for _, card := range game.Players[oldCurrentPlayer].Cards {
		if card.Color == cardToPlay.Color && card.Value == cardToPlay.Value {
			assert.Fail(t, "Wrong card was removed from the players hand.")
		}
	}
}

func TestPlayCard_DrawTwo(t *testing.T) {

	database, err := db.GetDb()

	game := setUpTestPlayerCard()

	game.CurrentPlayer = 3
	database.SaveGame(*game)

	oldCurrentPlayer := 3
	oldDrawPileSize := len(game.DrawPile)

	// Play a Draw Two from player4 hand
	cardToPlay := model.Card{Color: "red", Value: "D2"}
	game, err = playCard(game.ID, game.Players[3].ID, cardToPlay)

	// Assert we did not get an error.
	// Assert discard pile has a new card
	// Assert top card is card played
	// Assert player4 has 1 less card
	// Assert current player was shifted by one, so it should be player1 turn
	// Assert proper card was removed from players hand
	// Assert Player1 received two cards
	// Assert Draw pile lost 2 cards
	assert.Nil(t, err)
	assert.Equal(t, 2, len(game.DiscardPile))
	assert.Equal(t, cardToPlay.Color, game.DiscardPile[len(game.DiscardPile)-1].Color)
	assert.Equal(t, cardToPlay.Value, game.DiscardPile[len(game.DiscardPile)-1].Value)
	assert.Equal(t, 6, len(game.Players[oldCurrentPlayer].Cards))
	assert.Equal(t, 0, game.CurrentPlayer)
	for _, card := range game.Players[oldCurrentPlayer].Cards {
		if card.Color == cardToPlay.Color && card.Value == cardToPlay.Value {
			assert.Fail(t, "Wrong card was removed from the players hand.")
		}
	}
	assert.Equal(t, 9, len(game.Players[0].Cards))
	assert.Equal(t, oldDrawPileSize-2, len(game.DrawPile))

}

func TestPlayCard_Reverse(t *testing.T) {

	var err error
	game := setUpTestPlayerCard()

	//Set card to play that exists in player1 hand
	cardToPlay := model.Card{Color: "red", Value: "R"}
	oldGameDirection := game.Direction
	oldCurrentPlayer := 0

	// Play a reverse card
	game, err = playCard(game.ID, game.Players[0].ID, cardToPlay)

	// Assert we did not get an error.
	// Assert discard pile has a new card
	// Assert top card is card played
	// Assert player1 has 1 less card
	// Assert current player was shifted backwards by one so it's player 4
	// Assert proper card was removed from players hand
	// Assert Game Direction is not the same
	assert.Nil(t, err)
	assert.Equal(t, 2, len(game.DiscardPile))
	assert.Equal(t, cardToPlay.Color, game.DiscardPile[len(game.DiscardPile)-1].Color)
	assert.Equal(t, cardToPlay.Value, game.DiscardPile[len(game.DiscardPile)-1].Value)
	assert.Equal(t, 6, len(game.Players[oldCurrentPlayer].Cards))
	assert.Equal(t, 3, game.CurrentPlayer)
	for _, card := range game.Players[oldCurrentPlayer].Cards {
		if card.Color == cardToPlay.Color && card.Value == cardToPlay.Value {
			assert.Fail(t, "Wrong card was removed from the players hand.")
		}
	}
	assert.NotEqual(t, oldGameDirection, game.Direction)
}

func TestPlayCard_Wild(t *testing.T) {

	database, err := db.GetDb()

	game := setUpTestPlayerCard()

	game.CurrentPlayer = 3
	oldCurrentPlayer := 3
	database.SaveGame(*game)

	//Set card to play that exists in player4 hand
	cardToPlay := model.Card{Color: "blue", Value: "W"}

	// Play a wild card  setting it to blue
	game, err = playCard(game.ID, game.Players[3].ID, cardToPlay)

	// Assert we did not get an error.
	// Assert discard pile has a new card
	// Assert top card is card played
	// Assert player4 has 1 less card
	// Assert current player was shifted up by one
	// Assert proper card was removed from players hand
	assert.Nil(t, err)
	assert.Equal(t, 2, len(game.DiscardPile))
	assert.Equal(t, cardToPlay.Color, game.DiscardPile[len(game.DiscardPile)-1].Color)
	assert.Equal(t, cardToPlay.Value, game.DiscardPile[len(game.DiscardPile)-1].Value)
	assert.Equal(t, 6, len(game.Players[oldCurrentPlayer].Cards))
	assert.Equal(t, 0, game.CurrentPlayer)
	for _, card := range game.Players[oldCurrentPlayer].Cards {
		if card.Color == cardToPlay.Color && card.Value == cardToPlay.Value {
			assert.Fail(t, "Wrong card was removed from the players hand.")
		}
	}
}

func TestPlayCard_WildFour(t *testing.T) {

	database, err := db.GetDb()

	game := setUpTestPlayerCard()

	game.CurrentPlayer = 2
	oldCurrentPlayer := 2
	database.SaveGame(*game)

	//Set card to play that exists in player3 hand
	cardToPlay := model.Card{Color: "green", Value: "W4"}
	oldPlayer4CardCount := len(game.Players[3].Cards)
	// Play a wild draw 4 setting it to green
	game, err = playCard(game.ID, game.Players[2].ID, cardToPlay)

	// Assert we did not get an error.
	// Assert discard pile has a new card
	// Assert top card is card played
	// Assert player3 has 1 less card
	// Assert current player was shifted up by one
	// Assert proper card was removed from players hand
	// Assert Player 4 got 4 extra cards
	assert.Nil(t, err)
	assert.Equal(t, 2, len(game.DiscardPile))
	assert.Equal(t, cardToPlay.Color, game.DiscardPile[len(game.DiscardPile)-1].Color)
	assert.Equal(t, cardToPlay.Value, game.DiscardPile[len(game.DiscardPile)-1].Value)
	assert.Equal(t, 6, len(game.Players[oldCurrentPlayer].Cards))
	assert.Equal(t, 3, game.CurrentPlayer)
	for _, card := range game.Players[oldCurrentPlayer].Cards {
		if card.Color == cardToPlay.Color && card.Value == cardToPlay.Value {
			assert.Fail(t, "Wrong card was removed from the players hand.")
		}
	}
	assert.Equal(t, oldPlayer4CardCount+4, len(game.Players[3].Cards))
}

func TestPlayCard_PlaySameValue(t *testing.T) {

	database, err := db.GetDb()

	game := setUpTestPlayerCard()

	game.CurrentPlayer = 1
	oldCurrentPlayer := 1
	database.SaveGame(*game)

	// Card known in player2 hand
	cardToPlay := model.Card{Color: "blue", Value: "0"}

	// Play card with same value as top card but not same color
	game, err = playCard(game.ID, game.Players[1].ID, cardToPlay)

	// Assert we did not get an error.
	// Assert discard pile has a new card
	// Assert top card is card played
	// Assert player2 has 1 less card
	// Assert current player was shifted up by one
	// Assert proper card was removed from players hand
	assert.Nil(t, err)
	assert.Equal(t, 2, len(game.DiscardPile))
	assert.Equal(t, cardToPlay.Color, game.DiscardPile[len(game.DiscardPile)-1].Color)
	assert.Equal(t, cardToPlay.Value, game.DiscardPile[len(game.DiscardPile)-1].Value)
	assert.Equal(t, 6, len(game.Players[oldCurrentPlayer].Cards))
	assert.Equal(t, 2, game.CurrentPlayer)
	for _, card := range game.Players[oldCurrentPlayer].Cards {
		if card.Color == cardToPlay.Color && card.Value == cardToPlay.Value {
			assert.Fail(t, "Wrong card was removed from the players hand.")
		}
	}

}

func TestDrawCard(t *testing.T) {

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

}

func TestDealCards(t *testing.T) {

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

}

func TestCheckForCardInHand(t *testing.T) {
	//Created two cards One will be in the hand and the other won't
	validCard := model.Card{Color: "red", Value: "1"}
	falseCard := model.Card{Color: "blue", Value: "4"}
	//Created a hand with the valid card in it
	hand := []model.Card{validCard}

	//Testing to see if the function returns True for a card that is
	//present and False for a card that isn't present
	assert.True(t, checkForCardInHand(validCard, hand))
	assert.False(t, checkForCardInHand(falseCard, hand))
}

func TestCreatePlayer(t *testing.T) {
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
}

func TestJoinGame(t *testing.T) {
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
}
