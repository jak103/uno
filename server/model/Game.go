package model

// GameStatus status for a game
type GameStatus string

// Possible game status
const (
	WaitingForPlayers GameStatus = "Waiting For Players"
	Playing           GameStatus = "Playing"
	Finished          GameStatus = "Finished"
)

// Game Provies full game state
type Game struct {
	ID            string     `bson:"_id,omitempty" json:"id"`
	Name          string     `bson:"name,omitempty" json:"name"`
	Creator       Player     `bson:"creator,omitempty" json:"creator"`
	Password      string     `bson:"password,omitempty" json:"password"`
	DrawPile      []Card     `bson:"draw_pile,omitempty" json:"draw_pile"`
	DiscardPile   []Card     `bson:"discard_pile,omitempty" json:"discard_pile"`
	Players       []Player   `bson:"players,omitempty" json:"players"`
	CurrentPlayer int        `bson:"current_player,omitempty" json:"current_player"`
	Status        GameStatus `bson:"status,omitempty" json:"status"`
	Direction     bool       `bson:"direction,omitempty" json:"direction"`
}

// GameSummary Provides summary information for the lobby
type GameSummary struct {
	ID      string     `json:"id"`
	Name    string     `json:"name"`
	Creator string     `json:"creator"`
	Players []string   `json:"players"`
	Status  GameStatus `json:"status"`
}

// GameToSummary Converts a Game to a GameSummary
func GameToSummary(g Game) (summary GameSummary) {
	summary.ID = g.ID
	summary.Name = g.Name
	summary.Creator = g.Creator.Name
	summary.Status = g.Status

	for _, p := range g.Players {
		summary.Players = append(summary.Players, p.Name)
	}

	return summary
}
