package dbase

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func ConnectWithRetry(maxRetries int, delay time.Duration) (*sql.DB, error) {
	var db *sql.DB
	var err error
	connStr := "postgres://arkonterrr:arkonterrr%40yandex.ru@db:5432/db?sslmode=disable"

	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Printf("Failed to open DB connection: %v", err)
		} else {
			err = db.Ping()
			if err == nil {
				log.Println("Successfully connected to DB")
				return db, nil
			}
			log.Printf("Failed to ping DB: %v", err)
		}

		log.Printf("Retrying to connect in %s... (%d/%d)", delay, i+1, maxRetries)
		time.Sleep(delay)
	}
	return nil, err
}

func InitSchema(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			name TEXT NOT NULL DEFAULT '',
			company TEXT NOT NULL DEFAULT '',
			created_at TIMESTAMPTZ DEFAULT now()
		);`,
		`CREATE TABLE IF NOT EXISTS items (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT,
			image VARCHAR(512) NOT NULL DEFAULT '',
			created_at TIMESTAMPTZ DEFAULT now(),
			updated_at TIMESTAMPTZ DEFAULT now()
		);`,
		// сюда можно добавлять другие таблицы
	}

	for _, q := range queries {
		if _, err := db.Exec(q); err != nil {
			return err
		}
	}

	return nil
}
