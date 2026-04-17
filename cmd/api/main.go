package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Game struct {
	AppID       int    `json:"appid"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/games", gamesHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func gamesHandler(w http.ResponseWriter, r *http.Request) {
	games := []Game{
		{AppID: 1, Name: "Game 1", Description: "Description for Game 1"},
		{AppID: 2, Name: "Game 2", Description: "Description for Game 2"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}
