package model

// Represents a Message in the Chat
type Notification struct {
	Value  string `bson:"notification,omitempty" json:"notification"`
}
