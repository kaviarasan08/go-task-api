package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgresDb() (*sql.DB, error){
	connStr := "postgres://postgres:1425@localhost:5432/go_task_api"

	ctx, cancel:= context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	db, err := sql.Open("pgx", connStr)
	
	if err!= nil {
		return nil, err
	}

	if err = db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("database Connection failde: %w", err)
	}

	return db, nil

}