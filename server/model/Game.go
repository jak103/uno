package model

type Game struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Password string `bson:"password,omitempty" json:"password"`
}
