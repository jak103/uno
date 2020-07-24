package main

import (
	"encoding/json"
	"fmt"

	"github.com/jak103/uno/db"
)

type HandInfo struct {
	PlayerName string `json:"name"`
	CardCount  int    `json:"numCards"`
}

func getCardCount(gameID string) (string, error) {

	var result []*HandInfo = []*HandInfo{}

	database, err := db.GetDb()
	if err != nil {
		return "", fmt.Errorf("Could not connect to database : %v", err)
	}

	gameState, err := database.LookupGameByID(gameID)
	if err != nil {
		return "", fmt.Errorf("Unable to retrieve game state : %v", err)
	}

	for _, player := range gameState.Players {
		result = append(result, &HandInfo{player.Name, len(player.Cards)})
	}

	jsonResult, _ := json.Marshal(result)

	return string(jsonResult), nil
}
