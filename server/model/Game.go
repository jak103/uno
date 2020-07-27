package model

// GameStatus status for a game
type GameStatus string

// Possible game status
const (
	WaitingForPlayers GameStatus = "WaitingForPlayers"
	Playing           GameStatus = "Playing"
	Finished          GameStatus = "Finished"
)

type Game struct {
	ID            string     `bson:"_id,omitempty" json:"id"`
	Password      string     `bson:"password,omitempty" json:"password"`
	DrawPile      []Card     `bson:"draw_pile,omitempty" json:"draw_pile"`
	DiscardPile   []Card     `bson:"discard_pile,omitempty" json:"discard_pile"`
	Players       []Player   `bson:"players,omitempty" json:"players"`
	CurrentPlayer int        `bson:"current_player,omitempty" json:"current_player"`
	Status        GameStatus `bson:"status,omitempty" json:"status"`
	Direction     bool       `bson:"direction,omitempty" json:"direction"`
}

// HandInfo is a JSON struct for holding summary information about each players hand for a given game.
type HandInfo struct {
	PlayerName string `json:"name"`
	CardCount  int    `json:"numCards"`
}
