package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type SearchRequest struct {
	Name   string `json:"query"`
	Page   int    `json:"page" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
}

type Item struct {
	ID          string  `db:"id" json:"id"` // UUID из БД
	Title       string  `db:"title" json:"title"`
	Description string  `db:"description" json:"description"`
	Image       string  `db:"image" json:"image"`
	Price       float64 `db:"price" json:"price"` // Если нужна цена
}

func SearchHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SearchRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if req.Page < 1 {
			req.Page = 1
		}
		if req.Amount < 1 {
			req.Amount = 10
		}

		offset := (req.Page - 1) * req.Amount

		var items []Item
		var err error

		if req.Name == "" {
			// Без фильтрации
			err = db.Select(&items, `
				SELECT id, title, description, image, price
				FROM items
				ORDER BY created_at DESC
				LIMIT $1 OFFSET $2
			`, req.Amount, offset)
		} else {
			// Фильтрация по title
			err = db.Select(&items, `
				SELECT id, title, description, image, price
				FROM items
				WHERE title ILIKE $1
				ORDER BY created_at DESC
				LIMIT $2 OFFSET $3
			`, "%"+req.Name+"%", req.Amount, offset)
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB query failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"results": items})
	}
}
