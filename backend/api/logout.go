package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req TokenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Так как серверная сессия не хранится — просто подтверждаем выход
		c.String(http.StatusOK, "Logged out")
	}
}
