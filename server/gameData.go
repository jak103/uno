package main

import (
	"encoding/json"
	"fmt"
)

type HandInfo struct {
	PlayerName string `json:"name"`
	CardCount  int    `json:"numCards"`
}

func getCardCount(gameID string) (string, error) {

	var result []*HandInfo = []*HandInfo{}

	players, err := getPlayersInGame(gameID)
	if err != nil {
		return "", fmt.Errorf("Could not get list of players in %s: %v", gameID, err)
	}
	for _, player := range players {
		playerHand, err := getHandForPlayerInGame(gameID, player)
		if err != nil {
			return "", fmt.Errorf("Could not get hand for player '%s' in game %s: %v", player, gameID, err)
		}

		result = append(result, &HandInfo{player, len(playerHand)})
	}

	jsonResult, _ := json.Marshal(result)

	return string(jsonResult), nil
}

func getPlayersInGame(gameID string) ([]string, error) {
	// TODO: Remove error test
	if gameID == "errortest" {
		return []string{"getPlayersInGame"}, fmt.Errorf("There was an error in getPlayersInGame")
	}

	// TODO: Remove dummy data
	return []string{"player 1", "player 2", "player 3"}, nil
}

func getHandForPlayerInGame(gameID string, playerID string) ([]Card, error) {
	// TODO: Remove error test
	if gameID == "errortest" {
		return []Card{Card{-1, "getHandForPlayerInGame"}}, fmt.Errorf("There was an error in getHandForPlayerInGame")
	}

	// TODO: Remove dummy data
	if playerID == "player 1" {
		return []Card{newRandomCard()[0], newRandomCard()[0], newRandomCard()[0]}, nil
	}
	if playerID == "player 2" {
		return []Card{newRandomCard()[0]}, nil
	}

	return []Card{}, nil
}
