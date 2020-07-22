package model

import "github.com/google/uuid"

type Game struct {
	ID       uuid.UUID // UUID of the game.
	Password string    // Join code or password for lobby.
}
