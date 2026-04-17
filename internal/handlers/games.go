package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wyattprutch/gameboxd/internal/steam"
)

// holds dependencies for game-related handlers
type GamesHandler struct {
	Steam *steam.Client
}

// creates a GamesHandler with its dependencies
func NewGamesHandler(steamClient *steam.Client) *GamesHandler {
	return &GamesHandler{Steam: steamClient}
}

// search handles GET /api/games/search?q=halo
func (h *GamesHandler) Search(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing query parameter: q"})
		return
	}

	games, err := h.Steam.SearchGames(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to search Steam"})
		return
	}

	c.JSON(http.StatusOK, games)
}
