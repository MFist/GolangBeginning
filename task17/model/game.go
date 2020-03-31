package model

// Game this is an entity of a game
type Game struct {
	Name    string `json:"gameName"`
	Genre   string `json:"genre"`
	Year    int    `json:"releaseYear"`
	Console string `json:"console"`
}
