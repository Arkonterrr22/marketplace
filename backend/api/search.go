package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchRequest struct {
	Name   string `json:"query"`
	Page   int    `json:"page" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
}

type Item struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func SearchHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SearchRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		fmt.Print(req)
		if req.Page < 1 {
			req.Page = 1
		}
		if req.Amount < 1 {
			req.Amount = 10
		}

		var (
			rows *sql.Rows
			err  error
		)

		offset := (req.Page - 1) * req.Amount

		if req.Name == "" {
			// Если name пустой, не фильтруем
			rows, err = db.Query(`
				SELECT id, title, description, image
				FROM items
				ORDER BY id DESC
				LIMIT $1 OFFSET $2
			`, req.Amount, offset)
		} else {
			// Если name есть, фильтруем по частичному совпадению
			rows, err = db.Query(`
				SELECT id, title, description, image
				FROM items
				WHERE title ILIKE $1
				ORDER BY id DESC
				LIMIT $2 OFFSET $3
			`, "%"+req.Name+"%", req.Amount, offset)
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB query failed"})
			return
		}
		defer rows.Close()

		items := []Item{}
		for rows.Next() {
			var item Item
			if err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.Image); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan DB row"})
				return
			}
			items = append(items, item)
		}

		c.JSON(http.StatusOK, gin.H{"results": items})
	}
}
