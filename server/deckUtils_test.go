package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jak103/uno/model"
	
	//Needed for captureOutput()
	"bytes"
	"io"
	"log"
	"os"
	"sync"
)

func TestGenerateShuffledDeck(t *testing.T) {
	deck := generateShuffledDeck()

	//Check that the deck has the right number of each color
	colorCounts := map[string]int {
		"red": 0,
		"blue": 0,
		"green": 0,
		"yellow": 0,
		"wild": 0,
	}

	for _, card := range deck {
		colorCounts[card.Color]++
	}
	
	assert.Equal(t, 25, colorCounts["red"])
	assert.Equal(t, 25, colorCounts["blue"])
	assert.Equal(t, 25, colorCounts["green"])
	assert.Equal(t, 25, colorCounts["yellow"])
	assert.Equal(t, 8, colorCounts["wild"])

	//Check that the deck has the right number of total cards
	assert.Equal(t, 108, len(deck))
}

func TestShuffleCards(t *testing.T) { 
	deck := shuffleCards([]model.Card{model.Card{"red", "1"},model.Card{"blue", "2"},model.Card{"green", "3"}}) //Shuffling test deck
	assert.NotEqual(t, deck[:0], model.Card{"red", "1"}) 
}

func TestPrintCard(t *testing.T) {
	card := model.Card{"red", "1"} //Card to print
	re := captureOutput(func() { //Capturing the output from printcard 
		printCard(card) 
	})
	assert.Equal(t,"red 1\n",re) 
}

func TestPrintCards(t *testing.T) { 
	deck := []model.Card{model.Card{"red", "1"}, model.Card{"blue", "2"}, model.Card{"green", "3"}} //Building test deck
	re := captureOutput(func() { //Capturing output from printcards
		printCards(deck) 
	})
	assert.Equal(t,"red 1\nblue 2\ngreen 3\n",re) 
}

func captureOutput(f func()) string { //Function to capture output source: https://medium.com/@hau12a1/golang-capturing-log-println-and-fmt-println-output-770209c791b4
	reader, writer, err := os.Pipe() 
	if err != nil { 
		panic(err)
	}
	stdout := os.Stdout //Storing output location for later
	stderr := os.Stderr //Storing error location for later
	defer func() { //This will reset the values once this function returns
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer //Redirecting the output
	os.Stderr = writer //Redirecting the error
	log.SetOutput(writer) //Redirecting the log output
	out := make(chan string) //Creating final output variable 
	wg := new(sync.WaitGroup) 
	wg.Add(1)
	go func() { //The "reader" and the "writer" must not exist in the same Go function. This function separates the two. 
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait() 
	f() 
	writer.Close()  
	return <-out 
}
