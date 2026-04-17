package models

type Game struct {
	AppID       int    `json:"appid"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
