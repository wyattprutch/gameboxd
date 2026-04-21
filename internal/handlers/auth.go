package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wyattprutch/gameboxd/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	DB        *pgxpool.Pool
	JWTSecret []byte
}

func NewAuthHandler(db *pgxpool.Pool, jwtSecret string) *AuthHandler {
	return &AuthHandler{
		DB:        db,
		JWTSecret: []byte(jwtSecret),
	}
}

// Register handles POST /api/auth/register
func (h *AuthHandler) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Never store plain text passwords
	// bcrypt hashes the password with a random salt
	// The 12 is the cost factor - higher is slower but more secure
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	var user models.User
	err = h.DB.QueryRow(
		context.Background(),
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, username, email, created_at",
		input.Username, input.Email, string(hashedPassword),
	).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "username or email already exists"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Login handles POST /api/auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err := h.DB.QueryRow(
		context.Background(),
		"SELECT id, username, email, password FROM users WHERE email = $1",
		input.Email,
	).Scan(&user.ID, &user.Username, &user.Email, &user.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(h.JWTSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user":  user,
	})
}
