package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func NewPostgresDb() (*sql.DB, error) {

	err := godotenv.Load()

	if err != nil {
		return nil, fmt.Errorf("Env Load Failed: %w", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPswd := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPswd, dbHost, dbPort, dbName)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	db, err := sql.Open("pgx", dbUrl)

	if err != nil {
		return nil, err
	}

	if err = db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("database Connection failde: %w", err)
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(time.Minute * 25)
	db.SetConnMaxIdleTime(5 * time.Minute)

	return db, nil

}
