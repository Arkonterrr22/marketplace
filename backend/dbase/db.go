package dbase

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectWithRetry(maxRetries int, delay time.Duration) (*sqlx.DB, error) {
	connStr := "postgres://arkonterrr:arkonterrr%40yandex.ru@db:5432/db?sslmode=disable"

	var db *sqlx.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		db, err = sqlx.Connect("postgres", connStr)
		if err == nil {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if pingErr := db.PingContext(ctx); pingErr == nil {
				log.Println("Successfully connected to DB")
				return db, nil
			} else {
				err = pingErr
				log.Printf("Failed to ping DB: %v", err)
			}
		} else {
			log.Printf("Failed to connect to DB: %v", err)
		}

		log.Printf("Retrying to connect in %s... (%d/%d)", delay, i+1, maxRetries)
		time.Sleep(delay)
	}

	return nil, err
}

func InitSchema(db *sqlx.DB) error {
	queries := []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
		`CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			username TEXT NOT NULL DEFAULT '',
			company TEXT NOT NULL DEFAULT '',
			inn TEXT NOT NULL DEFAULT '',
			created_at TIMESTAMPTZ DEFAULT now()
		);`,
		`CREATE TABLE IF NOT EXISTS items (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			owner_id UUID NOT NULL REFERENCES users(id),
			title VARCHAR(255) NOT NULL,
			description TEXT,
			price NUMERIC(10,2),
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
