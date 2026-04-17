package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wyattprutch/gameboxd/internal/models"
)

// GetGames handles GET requests to retrieve a list of games
func GetGames(c *gin.Context) {
	games := []models.Game{
		{AppID: 1, Name: "Game 1", Description: "Description for Game 1"},
		{AppID: 2, Name: "Game 2", Description: "Description for Game 2"},
	}

	c.JSON(http.StatusOK, games)
}

// GetGame handles GET /api/games:appid
func GetGame(c *gin.Context) {
	appid := c.Param("appid")

	c.JSON(http.StatusOK, gin.H{
		"appid": appid,
		"note":  "placeholder",
	})
	// Implementation for retrieving a single game by app ID
}
