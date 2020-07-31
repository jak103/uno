package model

type PlayCard struct {
	Card     Card   `json:"card"`
	NewColor string `json:"color"`
}

type CallUno struct {
	Player Player `json:"player"`
}
