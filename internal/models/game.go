package models

// game is a game in the system
type Game struct {
	AppID       int    `json:"appid"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
