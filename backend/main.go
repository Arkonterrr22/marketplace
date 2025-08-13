package main

import (
	"backend/api"
	"backend/dbase"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err := dbase.ConnectWithRetry(10, 3*time.Second)
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}
	defer db.Close()

	if err := dbase.InitSchema(db); err != nil {
		log.Fatalf("Failed to initialize DB schema: %v", err)
	}

	r := gin.Default()

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://frontend:4321"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Роуты
	r.POST("/register", api.RegisterHandler(db))
	r.POST("/login", api.LoginHandler(db))
	r.POST("/me", api.MeHandler())
	r.POST("/logout", api.LogoutHandler())

	fmt.Println("Server started on :8080")
	r.Run(":8080")
}
