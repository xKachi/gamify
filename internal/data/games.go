package data

import (
	"time"
)

type Game struct {
	ID          int64       `json:id`                    // Unique integer ID for the game
	ReleaseYear ReleaseYear `json:"year,omitempty"`      // Year of game release
	CreatedAt   time.Time   `json:"-"`                   // Timestamp for when the game is added to our database
	Title       string      `json:"title"`               // Game title
	Developer   string      `json:"developer,omitempty"` // Name of the game developer
	Genres      []string    `json:"genres,omitempty"`    // Slice of genres for the game (action, adventure, etc.)
	Platforms   []string    `json:"platforms"`           // Slice of platforms for the game (Nintendo,Wi u, etc.)
	Version     int32       `json:"version"`             // The version number starts at 1 and will be incremented each
	// time the movie information is updated
}
