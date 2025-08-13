package api

import (
	"backend/auth"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name"`
	Company  string `json:"company"`
}

func RegisterHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req UserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		err := auth.Register(db, req.Email, req.Password, req.Name, req.Company)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Registration failed"})
			return
		}

		c.String(http.StatusCreated, "Registered")
	}
}
