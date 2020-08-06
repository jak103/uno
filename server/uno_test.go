package main

import (
	"testing"
	"github.com/jak103/uno/db"
	"github.com/jak103/uno/model"
	"github.com/stretchr/testify/assert"
	"errors"
)



// This function is meant to get a game and a player into the data base in a usable state for testing.
func setupGameWithPlayer(database *db.DB) (*model.Game, *model.Player) {
	player, _ := database.CreatePlayer("Player 1")

	game, _ := database.CreateGame("Game 1", player.ID)

	game.DrawPile = generateShuffledDeck(1)

	database.SaveGame(*game)

	return game, player
}

func TestCallUno(t *testing.T){

	//create two players and game place them in game
	database, _ := db.GetDb()
	unoPlayer,_ := database.CreatePlayer("UnoPlayer")
	callingPlayer,_ := database.CreatePlayer("CallingPlayer")
	game, unoPlayer := setupGameWithPlayer(database)

	game, _ = database.JoinGame(game.ID, unoPlayer.ID)
	game, _ = database.JoinGame(game.ID, callingPlayer.ID)

	//Deal cards to each player
	game, err := dealCards(game)
	assert.Nil(t, err, "error found")

	//getting rid of cards from uno player
	game.Players[0].Cards = game.Players[0].Cards[:1]
	
	//Call uno on player with one card, and no protection
	game, err1 := logicCallUno(game.ID, callingPlayer.ID, unoPlayer.ID)

	//Expect player to recieve four cards
	assert.Nil(t, err1, "error found")
	assert.Equal(t, 5, len(game.Players[0].Cards))

	//Get rid of cards from uno player to have one card
	game.Players[0].Cards = game.Players[0].Cards[:1]
	game.Players[0].Protection = true

	//Call uno on player with one card, and protection
	game, err2 := logicCallUno(game.ID, callingPlayer.ID, unoPlayer.ID)

	//Expect player to not receive cards
	assert.Nil(t, err2, "error found")
	assert.Equal(t, 1, len(game.Players[0].Cards))
}

func TestDrawCard(t *testing.T) {
	// Test passing in a bogus game id, we should get an error
	game, err := drawCard("Bogus game id", "Bogus player id")

	// Assert that we got an actual err
	assert.NotNil(t, err, "We did not error on a bogus game id")

	// Generate real game in database and real player
	database, _ := db.GetDb()
	game, player := setupGameWithPlayer(database)
	
	// Put a number card on the discard pile
	// For the purposes of this test, it's ok that it's an extra card
	game.DiscardPile = append(game.DiscardPile, model.Card{"red", "2"})
	database.SaveGame(*game)

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
	assert.Equal(t, 106, len(game.DrawPile))
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
	assert.NotNil(t, err, "Player drew out of turn. Please make sure only the player who's turn it is can play.")
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

func TestJoinGame(t *testing.T){
	// Get database
	database, err := db.GetDb()
	assert.Nil(t, err, "could not find database")
	// Create a new game with one player
	player, err := database.CreatePlayer("testPlayer")
	assert.Nil(t, err, "could not create new player")
	game, err := database.CreateGame("testGame", player.ID)
	assert.Nil(t, err, "could not create game")
	// Create a new player
	newPlayer, err := database.CreatePlayer("joinGamePlayer")
	assert.Nil(t, err, "could not create new player")
	// Attempt to join game
	game, err = joinGame(game.ID, newPlayer)
	database.SaveGame(*game)
	assert.Nil(t, err, "could not join game with new player")
	// Lookup game from database 
	game, err = database.LookupGameByID(game.ID)
	assert.Nil(t, err, "could not find game in database")
	// Test to see if the newPlayer is in the game
	assert.Contains(t, game.Players, *newPlayer)
	// attempt to join an errored game
	err = errors.New("MockDB: Error!")
	game, err = joinGame("Bad ID", newPlayer)
	assert.Nil(t, game, "Joined a valid game")

}

func TestDrawTopCard(t *testing.T) {
	// Creating database and testing for errors
	database, err := db.GetDb()
	assert.Nil(t, err, "MockDB: Could not retrive database")
	// Creating player and testing for errors
	player , err := database.CreatePlayer("Test Player")
	assert.Nil(t, err, "MockDB: Could not create player")
	// Creating game and testing for errors 
	game, err := database.CreateGame("Test Game", player.ID)
	assert.Nil(t, err, "MockDB: Could not create game")
	// Setting game.DrawPile to a test deck
	game.DrawPile = []model.Card{model.Card{"red", "1"}, model.Card{"blue", "2"}, model.Card{"green", "3"}}
	// Testing drawTopCard
	game, cardReturned := drawTopCard(game) 
	assert.Equal(t, model.Card{"green", "3"}, cardReturned)
}

func TestGoToNextPlayer(t *testing.T) {
	// Creating database and testing for errors
	database, err := db.GetDb()
	assert.Nil(t, err, "MockDB: Could not retrive database")
	// Creating creator and first player and testing for errors
	creator , err := database.CreatePlayer("Creator")
	player2 , err := database.CreatePlayer("Test 2")
	assert.Nil(t, err, "MockDB: Could not create player")
	// Creating third player to add to game and testing for errors
	player3 , err := database.CreatePlayer("Test 3")
	assert.Nil(t, err, "MockDB: Could not create player")
	// Creating game and testing for errors 
	game, err := database.CreateGame("Test Game 1", creator.ID)
	assert.Nil(t, err, "MockDB: Could not create game")
	// Adding players
	game , err  = joinGame(game.ID, player2)
	database.SaveGame(*game)
	game , err  = joinGame(game.ID, player3)
	database.SaveGame(*game)
	// Testing a situation where the players have no cards to trigger winning condition if statement.  
	game.CurrentPlayer = 0
	game = goToNextPlayer(game)
	// When winning condition is present goToNextPlayer will not change the current player
	assert.Equal(t, 0, game.CurrentPlayer)
	// Dealing cards to players
	game, err = dealCards(game)
	game.CurrentPlayer = 1
	assert.Nil(t, err, "MockDB: Could not deal cards")
	// Testing one direction
	game = goToNextPlayer(game)
	assert.Equal(t,2,game.CurrentPlayer)
	// Swapping direction
	game.Direction = !game.Direction
	// Testing the other direction 
	game = goToNextPlayer(game)
	assert.Equal(t,1,game.CurrentPlayer)
	game = goToNextPlayer(game)
	assert.Equal(t,0,game.CurrentPlayer)
	database.DeleteGame(game.ID)
}
func TestIsCardPlayable(t *testing.T){

	// Generate real game in database and real player
	database, err := db.GetDb()
	assert.Nil(t, err, "could not find database")

	game, _ := setupGameWithPlayer(database)

	// puts a card to test with in the discard pile
	game.DiscardPile = append(game.DiscardPile, model.Card{Color: "red", Value: "2"})

	// tests to see if a card of the same color is playable
	test1 := isCardPlayable(model.Card{Color: "red", Value: "1"} , game.DiscardPile)
	assert.Equal(t, test1, true)

	// tests to see if a card of the same number is playable
	test2 := isCardPlayable(model.Card{Color: "blue", Value: "2"} , game.DiscardPile)
	assert.Equal(t, test2, true)

	// tests to see if a wild is playable
	test3 := isCardPlayable(model.Card{Color: "black", Value: "W"} , game.DiscardPile)
	assert.Equal(t, test3, true)

	// tests to see if a wild draw four is playable
	test4 := isCardPlayable(model.Card{Color: "black", Value: "W4"} , game.DiscardPile)
	assert.Equal(t, test4, true)
}
func TestReshuffleDiscardPile(t *testing.T){

	// Generate real game in database and real player
	database, err := db.GetDb()
	assert.Nil(t, err, "could not find database")

	game, _ := setupGameWithPlayer(database)

	// puts the deck into the discard pile from the beginning
	game.DiscardPile = generateShuffledDeck(1)

	// shuffles the discard pile into the draw pile
	game = reshuffleDiscardPile(game)

	// checks to see if the discard pile is now empty
	assert.Equal(t, len(game.DiscardPile), 1)

}

func TestAddMessage(t *testing.T){
	
	// get database
	database, err := db.GetDb()
	assert.Nil(t, err, "could not find database")

	// sets up the game with a player, creates a message for the only player in the game
	game, player := setupGameWithPlayer(database)
	m := model.Message{Player: *player, Value: "Hello World"}
	
	// add the Message to the game
	game, err = addMessage(game.ID, player.ID, m)

	// test to see if the new Message is in the game
	assert.Contains(t, game.Messages, m)
}

func TestDrawNCards(t *testing.T){
	// get database
	database, err := db.GetDb()
	assert.Nil(t, err, "could not find database")
	// create game
	game, _, err := createNewGame("testGame", "testCreater")
	assert.Nil(t, err, "game not created")
	game.DrawPile = generateShuffledDeck(1)
	// create player
	player, err:= createPlayer("player2")
	assert.Nil(t, err, "player not created")
	// add player to game
	game, _ = joinGame(game.ID, player)
	database.SaveGame(*game)
	assert.Nil(t, err, "game not joined")
	// deal cards out
	game, err = dealCards(game)
	assert.Nil(t, err, "cards not dealt")
	// check if player 2 joined the game
	// players := len(game.Players)
	// assert.Equal(t, players, 2)
	// check if player was dealt cards
	assert.Equal(t, 7, len(game.Players[game.CurrentPlayer].Cards)) 
	// draw 2
	game = drawNCards(game, 2) 
	// check if 2 cards were drawn
	assert.Equal(t, 9, len(game.Players[game.CurrentPlayer].Cards)) 
	// draw 4
	game = drawNCards(game, 4) 
	// check if 4 cards were drawn
	assert.Equal(t, 13, len(game.Players[game.CurrentPlayer].Cards)) 
} 

func TestCheckGameExists(t *testing.T){
	// Get database
	database, err := db.GetDb()
	assert.Nil(t, err, "could not find database")
	// Create Player
	player, err := database.CreatePlayer("testPlayer")
	assert.Nil(t, err, "could not create player")
	// Create Game
	game, err := database.CreateGame("testGame", player.ID)
	assert.Nil(t, err, "could not create game")
	// Check to see if the function detects the created game
	validGame, err := checkGameExists(game.ID)
	assert.True(t, validGame)
	// Check to see if the function does not detect a game that does not exist
	fakeGame, err := checkGameExists("fakeGame")
	assert.False(t, fakeGame)
}

func TestGetGameUpdate(t *testing.T){
	// Get database
	database, err := db.GetDb()
	assert.Nil(t, err, "could not find database")
	// Create Player
	player, err := database.CreatePlayer("testPlayer")
	assert.Nil(t, err, "could not create player")
	// Create Game
	game, err := database.CreateGame("testGame", player.ID)
	assert.Nil(t, err, "could not create game")
	// Get Game Update from function
	gameUpdate, err := getGameUpdate(game.ID, player.ID)
	assert.Nil(t, err, "could not get game update")
	// Get game data from the database
	gameData, err := database.LookupGameByID(game.ID)
	assert.Nil(t, err, "could not get game from database")
	// Check to see if the gameUpdate is equal to the game in the database
	assert.Equal(t, gameData, gameUpdate)
	// Check that the function returns Nil for non existant game
	fakeGame, _ := getGameUpdate("fakeGame", "fakePlayer")
	assert.Nil(t, fakeGame, "Found game that does not exist")
}
