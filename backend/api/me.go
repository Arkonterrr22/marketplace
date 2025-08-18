package api

import (
	"backend/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Token string `json:"token" binding:"required"`
}

func MeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req TokenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		userToken, err := auth.ParseJWT(req.Token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user_id":  userToken.ID,
			"email":    userToken.Email,
			"username": userToken.Username,
			"company":  userToken.Company,
		})
	}
}
