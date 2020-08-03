package service

//"testing"

//"github.com/jak103/uno/db"
//"github.com/jak103/uno/model"

// This function is meant to get a game and a player into the data base in a usable state for testing.
//func setupGameWithPlayer(database *db.DB) (*model.Game, *model.Player) {

// game, _ := database.CreateGame()

// player, _ := database.CreatePlayer("Player 1")

// database.JoinGame(game.ID, player.ID)

// game.DrawPile = generateShuffledDeck()

// database.SaveGame(*game)
// return game, player
//}

//func TestDrawCard(t *testing.T) {

// database, _ := db.GetDb()

// game, player := setupGameWithPlayer(database)

// if drawCard(game.ID, player.ID) {
// 	player, _ = database.LookupPlayer(player.ID)
// 	assert.Equal(t, len(player.Cards), 1)
// 	assert.Equal(t, len(game.DrawPile), 107)
// } else {
// 	assert.Fail(t, "Failed to draw card.")
// }
//}
