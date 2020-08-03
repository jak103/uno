package model

// Represents a Message in the Chat
type Message struct {
	Player Player `bson:"player,omitempty" json:"player"`
	Value  string `bson:"message,omitempty" json:"message"`
}
