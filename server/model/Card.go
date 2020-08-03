package model

// Represents a card
// Uses Value instead of Number
// to more accurately represent non-numerical cards
type Card struct {
	Color string `bson:"color,omitempty" json:"color"`
	Value string `bson:"value,omitempty" json:"value"`
}
