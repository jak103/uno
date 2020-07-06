package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	
	//"fmt"
	
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestRandColor(t *testing.T){
	// check that each color is properly generated...
	assert.Equal(t, "red", randColor(0))
	assert.Equal(t, "blue", randColor(1))
	assert.Equal(t, "green", randColor(2))
	assert.Equal(t, "yellow", randColor(3))
	assert.Equal(t, "", randColor(4))
}

func TestNewRandomCard(t *testing.T){
	// get a thousand cards, check the ratios...
	colorsMap := make(map[string]int) 
	numsMap   := make(map[int]int)
	colorsMap["red"] = 0.0
	colorsMap["green"] = 0.0
	colorsMap["blue"] = 0.0
	colorsMap["yellow"] = 0.0
	colorsMap[""] = 0.0
	
	colors := []string{"red", "blue", "green", "yellow"}
	maxNum := 9
	minNum := 0
	
	for i := -1; i < 12; i++ {
		numsMap[i] = 0.0
	}
	
	totalCards := 1000
	
	
	// keep track of cards that aren't wild.
	// they don't exist yet, but will!
	// TODO: We can add checks for other types of cards in the future.
	coloredCards := 0
	
	// as opposed to skips, wilds, etc.
	// TODO: handle skips, wilds, etc.
	numberedCards := 0
	
	for i := 0; i < totalCards; i++ {
		testCard := newRandomCard()[0]
		
		// count the card if the card has a valid color... (not wild)
		if stringInArray(colors, testCard.Color){
			coloredCards++
		}
		
		// count the card if the card has a number (not a skip, etc)
		if testCard.Number <= maxNum && testCard.Number >= minNum{
			numberedCards++
		}
		
		// keep track of the amount of each color/number so we can check the ratio
		colorsMap[testCard.Color]++
		numsMap[testCard.Number]++
		
	}
	
	// we shouldn't have gotten any cards with these values.
	assert.Equal(t, colorsMap[""], 0)
	assert.Equal(t, numsMap[-1], 0)
	assert.Equal(t, numsMap[11], 0)
	
	// these values should be within a reasonable range for a thousand cards.
	// note: this is random, but it will fail a correct implementation less than .05% of the time
	
	// check that each of the ten numbers has about 10% (within 6%)
	for i := 1; i < 10; i++ {
		assert.InDelta(t, float64(numsMap[i])/float64(numberedCards), .1, 0.06)
	}
	
	// check that each of the 4 colors has about 25% (within 6%)
	assert.InDelta(t, float64(colorsMap["red"])/float64(coloredCards), .25, 0.6)
	assert.InDelta(t, float64(colorsMap["green"])/float64(coloredCards), .25, 0.06)
	assert.InDelta(t, float64(colorsMap["blue"])/float64(coloredCards), .25, 0.06)
	assert.InDelta(t, float64(colorsMap["yellow"])/float64(coloredCards), .25, 0.06)
	
}

func TestDealCards(t *testing.T){
	// make sure our game is a blank slate.
	players = make([]string, 0)
	allCards = make(map[string][]Card, 0)
	gameStarted = false
	
	// set up a fake game.
	players = append(players, "Bob", "Jill", "Chanel")
	
	dealCards()
	
	// check that each player got added to the all cards map.
	for _, player := range players {
		assert.Contains(t, allCards, player)
	}
	
	// make sure that each player has the same number of cards as Bob.
	for _, player := range players {
		assert.Equal(t, len(allCards["Bob"]), len(allCards[player]))
	}
	
	// check that a player was made the first player...
	assert.Contains(t, players, currPlayer)
	
	// make sure the game has started
	assert.Equal(t, true, gameStarted)
	
	// clear this, so it doesn't affect other tests.
	players = make([]string, 0)
	allCards = make(map[string][]Card, 0)
}

func TestCheckForWinner(t *testing.T){
	// make sure our game is a blank slate.
	players = make([]string, 0)
	allCards = make(map[string][]Card, 0)
	
	// set up a fake game.
	players = append(players, "Bob", "Jill", "Chanel")
	
	dealCards()
	
	// we just dealt cards, There shouldn't be a winner.
	assert.Equal(t, "", checkForWinner())
	
	// pretend that Jill played all her cards...
	allCards["Jill"] = make([]Card, 0)
	
	// Jill has no cards. She should be the winner!
	assert.Equal(t, "Jill", checkForWinner())
	
	// clear this, so it doesn't affect other tests.
	players = make([]string, 0)
	allCards = make(map[string][]Card, 0)
	
}

func TestPlayCard(t *testing.T){
	// make sure our game is a blank slate.
	players = make([]string, 0)
	allCards = make(map[string][]Card, 0)
	gameID = "1234"
	
	// set up a fake game.
	players = append(players, "Bob", "Jill", "Chanel")
	
	dealCards()
	
	// looked here for some setup help:
	// https://echo.labstack.com/guide/testing
	
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	testPlayerIndex := playerIndex // get started at the same point...
	
	// make a request with the current player
	c.SetParamNames("game", "username")
	c.SetParamValues("1234", players[testPlayerIndex])
	
	// loop through each color
	for i := 0; i < 4; i++ {
		// make a card of the right number, but not necessarily the right color
		// note: the color will ALWAYS be wrong after the first iteration.
		cardToPlay := Card{ currCard[0].Number,  randColor(i)}
		
		// give the current player the card we just created
		allCards[players[testPlayerIndex]] = append(allCards[players[testPlayerIndex]], cardToPlay)
		
		// get the current length of the players hand
		cardsLeft := len(allCards[players[testPlayerIndex]])
		
		// try playing same num
		r := playCard(c, cardToPlay)
		
		// did we get a valid response?
		assert.Equal(t, true, r.ValidGame)
		
		// did our card become the current card?
		assert.Equal(t, currCard[0], cardToPlay)
		
		// did our card leave our hand?
		assert.Equal(t, cardsLeft - 1, len(allCards[players[testPlayerIndex]]))
		
		// the player index should have shifted by one
		testPlayerIndex = (testPlayerIndex + 1) % len(players)
		assert.Equal(t, playerIndex, testPlayerIndex)
		
		// set up the request to be from the next player
		c.SetParamValues("1234", players[testPlayerIndex])
		
		// get a different number than the current card
		nextNum := (currCard[0].Number + 1) % 10		
		
		// create a card of the same color but different number as the current card
		cardToPlay = Card{ nextNum,  randColor(i)}
		
		// add this card to the current players hand
		allCards[players[testPlayerIndex]] = append(allCards[players[testPlayerIndex]], cardToPlay)
		
		// get the size of the hand.
		cardsLeft = len(allCards[players[testPlayerIndex]])
		
		// try playing same color
		r = playCard(c, cardToPlay)
		
		// did we get a valid response?
		assert.Equal(t, true, r.ValidGame)
		
		// did our card become the "current card"?
		assert.Equal(t, currCard[0], cardToPlay)
		
		// did our hand size decrease by exactly one?
		assert.Equal(t, cardsLeft - 1, len(allCards[players[testPlayerIndex]]))
		
		// did the player index move?
		testPlayerIndex = (testPlayerIndex + 1) % len(players)
		assert.Equal(t, playerIndex, testPlayerIndex)
		
		// set up a request with the next player
		c.SetParamValues("1234", players[testPlayerIndex])
		
		// get a number different from the current card...
		nextNum = (currCard[0].Number + 1) % 10
		
		// try all the wrong colors...
		for k := 1; k < 4; k++ {
			
			colorIndex := (i + k) % 4
			
			// will not match in color or number.
			cardToPlay = Card{ nextNum,  randColor(colorIndex)}
			allCards[players[testPlayerIndex]] = append(allCards[players[testPlayerIndex]], cardToPlay)
			
			cardsLeft = len(allCards[players[testPlayerIndex]])
			
			// try playing an invalid card!
			r := playCard(c, cardToPlay)
			
			// check that we got a valid response
			assert.Equal(t, true, r.ValidGame)
			// but the player index did not shift
			assert.Equal(t, playerIndex, testPlayerIndex)
			// the size of our hand did not change
			assert.Equal(t, cardsLeft, len(allCards[players[testPlayerIndex]]))
			// the card we attempted to play is NOT now the current card.
			assert.NotEqual(t, currCard[0], cardToPlay)
		}
		
	}	
	
	// if we don't check for this, people could try to hack and skip their turn...
	// try playing a valid card that we don't have...
	allCards[players[testPlayerIndex]] = make([]Card, 0)
	
	nextNum := (currCard[0].Number + 1) % 10
	
	nextColor := "red"
	
	if currCard[0].Color == nextColor {
		nextColor = "blue";
	}
	
	allCards[players[testPlayerIndex]] = append(allCards[players[testPlayerIndex]], Card{nextNum, nextColor}, Card{nextNum, nextColor})
	
	cardToPlay := Card{currCard[0].Number, nextColor}
	
	// while the card is valid, we don't have it.
	// check that the card doesn't play, and the turn doesn't change.
	r := playCard(c, currCard[0])
	assert.Equal(t, true, r.ValidGame)
	assert.NotEqual(t, currCard[0], cardToPlay)
	assert.Equal(t, playerIndex, testPlayerIndex)
	
	// try playing a valid card from the wrong game
	
	allCards[players[testPlayerIndex]] = append(allCards[players[testPlayerIndex]], cardToPlay)
	
	c.SetParamValues("12345", players[testPlayerIndex])
	
	r = playCard(c, currCard[0])
	assert.Equal(t, false, r.ValidGame)
	assert.NotEqual(t, currCard[0], cardToPlay)
	assert.Equal(t, playerIndex, testPlayerIndex)
	// check that the player was valid, but game was not (false response due to game)
	assert.Equal(t, currPlayer, players[testPlayerIndex])
	assert.Equal(t, false, checkID(c.Param("game")))
	
	// try playing a valid card as the wrong player
	
	testPlayerIndex = (testPlayerIndex + 1) % len(players)
	c.SetParamValues("1234", players[testPlayerIndex])
	
	r = playCard(c, currCard[0])
	assert.Equal(t, false, r.ValidGame)
	assert.NotEqual(t, currCard[0], cardToPlay)
	// check that the game was valid, but player was not (false response due to player)
	assert.Equal(t, true, checkID(c.Param("game")))
	assert.NotEqual(t, currPlayer, players[testPlayerIndex])
	
	// check to make sure the player index did not change.
	testPlayerIndex = testPlayerIndex - 1
	
	if testPlayerIndex < 0 {
		testPlayerIndex = len(players) - 1
	}
	
	assert.Equal(t, playerIndex, testPlayerIndex)
	
	// clear this, so it doesn't affect other tests.
	currPlayer = ""
	gameID = ""
	players = make([]string, 0)
	allCards = make(map[string][]Card, 0)
}

func TestCheckForCardInHand (t *testing.T) {
	// make sure our game is a blank slate.
	players = make([]string, 0)
	allCards = make(map[string][]Card, 0)
	gameID = "1234"
	
	// set up a fake game.
	players = append(players, "Bob", "Jill", "Chanel")
	
	dealCards()
	
	// clear Jill's hand
	allCards["Jill"] = make([]Card, 0)
	
	card := Card{5, "red"}
	
	// check to make sure empty hand returns false.
	assert.Equal(t, false, checkForCardInHand(card, "Jill"))
	
	// Same color, different number
	allCards["Jill"] = append(allCards["Jill"], Card{6, "red"})
	assert.Equal(t, false, checkForCardInHand(card, "Jill"))
	
	// Same number, different color
	allCards["Jill"] = append(allCards["Jill"], Card{5, "blue"})
	assert.Equal(t, false, checkForCardInHand(card, "Jill"))
	
	// Same card, different player
	allCards["Bob"] = append(allCards["Bob"], Card{5, "red"})
	assert.Equal(t, false, checkForCardInHand(card, "Jill"))
	
	// add the right card, make sure it is in the hand.
	allCards["Jill"] = append(allCards["Jill"], Card{5, "red"})
	assert.Equal(t, true, checkForCardInHand(card, "Jill"))
	
}

func TestContains (t *testing.T) {
	testArr := make([]string, 0)
	
	// empty array doesn't contain "a"
	index, ok := contains(testArr, "a")
	assert.Equal(t, false, ok)
	assert.Equal(t, -1, index)
	
	testArr = append(testArr, "a", "b", "c")
	
	// array doesn't contain "d"
	index, ok = contains(testArr, "d")
	assert.Equal(t, false, ok)
	assert.Equal(t, -1, index)
	
	// array doesn't contain "a"; index is 0
	index, ok = contains(testArr, "a")
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, index)
	
	// array doesn't contain "b"; index is 1
	index, ok = contains(testArr, "b")
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, index)
	
	// array doesn't contain "c"; index is 2
	index, ok = contains(testArr, "c")
	assert.Equal(t, true, ok)
	assert.Equal(t, 2, index)
	
}

func stringInArray(arr []string, el string) bool{
	
	for _, item := range arr {
		if item == el {
			return true
		}
	}
	
	return false
	
}