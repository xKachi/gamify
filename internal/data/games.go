package data

import (
	"time"

	"github.com/xKachi/gamify.git/internal/validator"
)

type Game struct {
	ID          int64     `json:id`
	CreatedAt   time.Time `json:"-"` 
	Title       string    `json:"title"`
	Playtime    Playtime  `json:"playtime,omitempty"`
	ReleaseYear int32     `json:"release_year,omitempty"` 
	Developer   string    `json:"developer,omitempty"`    // Name of the game developer
	Genres      []string  `json:"genres,omitempty"`       // Slice of genres for the game (action, adventure, etc.)
	Platforms   []string  `json:"platforms"`              // Slice of platforms for the game (Nintendo,Wi u, etc.)
	Version     int32     `json:"version"`                // The version number starts at 1 and will be incremented each
	
}

func ValidateGame(v *validator.Validator, game *Game) {
	v.Check(game.Title != "", "title", "must be provided")
	v.Check(len(game.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(game.ReleaseYear != 0, "release-year", "must be provided")
	v.Check(game.ReleaseYear >= 2000, "release-year", "must be greater than 2000")
	v.Check(game.ReleaseYear <= int32(time.Now().Year()), "release-year", "must not be in the future")

	v.Check(game.Playtime != 0, "playtime", "must be provided")
	v.Check(game.Playtime > 0, "playtime", "must be a positive integer")

	v.Check(game.Genres != nil, "genres", "must be provided")
	v.Check(len(game.Genres) >= 1, "genres", "must contain atleast one genres")
	v.Check(len(game.Genres) <= 3, "genres", "must not contain more than 3 genres")
	// Note that we're using the Unique helper in the line below to check that all
	// values in the game.Genres slice are unique.
	v.Check(validator.Unique(game.Genres), "genres", "must not contain duplicate values")

	v.Check(game.Platforms != nil, "platforms", "must be provided")
	v.Check(len(game.Platforms) >= 1, "platforms", "must contain atleast one platform")
	v.Check(len(game.Platforms) >= 3, "platforms", "must not contain more than 3 platforms")
	v.Check(validator.Unique(game.Platforms), "platforms", "must not contain duplicate values")
}
