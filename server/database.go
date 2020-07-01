package main

import(
	"context"
	"log"
	"cloud.google.com/go/firestore"

	// Debuggings
	"fmt"
)
// DB can be initialized 2 different ways
type DB struct {}

// Variables
var currCard []Card = nil // The cards are much easier to render as a list
var players []string = []string{}
var playerIndex = 0 // Used to iterate through the players
var currPlayer string = ""
var allCards map[string][]Card = make(map[string][]Card) // k: username, v: list of cards
var gameStarted bool = false
var gameID string = ""
var numberOfPlayers int = 0

func playerHasCard(arr []Card, val Card) (int, bool) {
	for i, item := range arr {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "usu-devops"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	defer client.Close()
	return client
}

//Setters
func (db *DB) addNewGame() string { // TODO: Should return an error if it doesn't work
	gameID = "12234" // TODO: Create random game id and add to firestore
	return gameID
}

func (db *DB) addUserToGame(username string) {
	if _, there := allCards[username]; !there {
		allCards[username] = nil; // No cards yet
		players = append(players, username)
		numberOfPlayers++;
	}

	fmt.Println("Number of players: ", numberOfPlayers)
}

func (db *DB) updateCards(cards map[string][]Card) {
	allCards = cards
}

func (db *DB) startGame(gameCode string) {
	gameStarted = true
	currPlayer = players[0]
}

func (db *DB) setCurrentCard(card []Card) {
	currCard = card
}

func (db *DB) removeCardFromPlayersDeck(card Card, gameCode, user string) {
	index, _ := playerHasCard(allCards[user], card)
	allCards[user][index] = allCards[user][len(allCards[user])-1]
	allCards[user][len(allCards[user])-1] = Card{}
	allCards[user] = allCards[user][:len(allCards[user])-1]
}

///////////////////////////////////////////////////////////////////////////////////////
// Getters. TODO: Return nil if gameCode is bad
///////////////////////////////////////////////////////////////////////////////////////
func (db *DB) getCurrentCard(gameCode string) []Card {
	return currCard // TODO: Make the call to firestore and return the actual values
}

func (db *DB) getAllPlayers(gameCode string) []string {
	return players // TODO: Make the call to firestore and return the actual values
}

func (db *DB) getAllCards(gameCode string) map[string][]Card {
	return allCards // TODO: Make the call to firestore and return the actual values
}

func (db *DB) hasGameStarted(gameCode string) bool {
	return gameStarted // TODO: Make the call to firestore and return the actual values
}

func (db *DB) getCurrentPlayer(gameCode string) string {
	return currPlayer // TODO: Make the call to firestore and return the actual values
}

func (db *DB) getCurrentPlayerCards(gameCode, username string) []Card {
	return allCards[username] // TODO: Make the call to firestore and return the actual values
}

func (db *DB) isValidGame(gameCode string) bool {
	return gameID == gameCode // TODO: Make the call to firestore and return the actual values
}

func (db *DB) isInDeck(card Card, gameCode, user string) bool {
	_, hasCard := playerHasCard(allCards[user], card)
	return hasCard 
}
