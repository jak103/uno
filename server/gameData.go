package main

import "fmt"

func getCardCount(gameID string) (string, error) {
	if gameID == "errortest" {
		return "invalid data", fmt.Errorf("There was an error")
	}
	return "this is a test", nil
}
