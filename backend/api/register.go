package api

import (
	"backend/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req auth.User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		user := auth.User{
			Email:    auth.Email(req.Email),
			Password: auth.PasswordHash(req.Password),
			Username: auth.Username(req.Username),
			Company:  auth.Company(req.Company),
		}

		err := auth.Register(c.Request.Context(), db, &user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Registration failed"})
			return
		}

		c.String(http.StatusCreated, "Registered")
	}
}
