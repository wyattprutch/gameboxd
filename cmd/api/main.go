package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wyattprutch/gameboxd/internal/config"
	"github.com/wyattprutch/gameboxd/internal/db"
	"github.com/wyattprutch/gameboxd/internal/handlers"
	"github.com/wyattprutch/gameboxd/internal/middleware"
	"github.com/wyattprutch/gameboxd/internal/steam"
)

func main() {
	// load .env file in development
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from environment")
	}

	cfg := config.Load()

	if cfg.SteamAPIKey == "" {
		log.Fatal("STEAM_API_KEY is required")
	}

	database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()
	log.Println("Connected to database")

	// create the Steam client and inject it into handlers
	steamClient := steam.NewClient(cfg.SteamAPIKey)
	gamesHandler := handlers.NewGamesHandler(steamClient)
	authHandler := handlers.NewAuthHandler(database, cfg.JWTSecret)

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	{
		api.POST("/auth/register", authHandler.Register)
		api.POST("/auth/login", authHandler.Login)
		api.GET("/games/search", gamesHandler.Search)

		// protected routes
		protected := api.Group("/")
		protected.Use(middleware.RequireAuth(cfg.JWTSecret))
		{
			protected.GET("/me", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"user_id":  c.MustGet("user_id"),
					"username": c.MustGet("username"),
				})
			})
		}

	}

	log.Printf("Server starting on port %s", cfg.Port)
	r.Run(":" + cfg.Port)
}
