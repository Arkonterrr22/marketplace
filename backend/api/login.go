package api

import (
	"backend/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func LoginHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req auth.User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		user, err := auth.Authenticate(c.Request.Context(), db, &req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials!", "creds": &req})
			return
		}

		tokenString, err := auth.GenerateJWT(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token":    tokenString,
			"email":    user.Email,
			"username": user.Username,
			"company":  user.Company,
		})
	}
}
