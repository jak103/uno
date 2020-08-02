package model

// Player Model that represents a Player and their hand
type Player struct {
	ID    string `bson:"_id,omitempty" json:"id"`
	Name  string `bson:"name,omitempty" json:"name"`
	Cards []Card `bson:"cards,omitempty" json:"cards"`
}
