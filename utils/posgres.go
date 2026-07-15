package utils

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func Conn() *pgx.Conn {
	connStr := os.Getenv("DB")
	var db *pgx.Conn
	db, _ = pgx.Connect(context.Background(), connStr)
	return db
}