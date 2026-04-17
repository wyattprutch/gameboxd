package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wyattprutch/gameboxd/internal/config"
	"github.com/wyattprutch/gameboxd/internal/handlers"
)

func main() {
	cfg := config.Load()

	// Initialize Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		// OPTIONS requests are used for CORS preflight checks, so can return early
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Route group
	api := r.Group("/api")
	{
		api.GET("/games", handlers.GetGames)
		api.GET("/games/:appid", handlers.GetGame)
	}

	log.Printf("Server starting on port %s", cfg.Port)
	r.Run(":" + cfg.Port)
}
