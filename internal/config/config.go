package config

import "os"

// holds all values loaded from environment variables
type Config struct {
	Port        string
	SteamAPIKey string
	JWTSecret   string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
}

// load reads environment variables and returns a Config struct
// if var isnt found, it uses the provided default value
func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		SteamAPIKey: getEnv("STEAM_API_KEY", ""),
		JWTSecret:   getEnv("JWT_SECRET", "supersecretkey"),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBUser:      getEnv("DB_USER", "gameboxd"),
		DBPassword:  getEnv("DB_PASSWORD", "password"),
		DBName:      getEnv("DB_NAME", "gameboxd"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
