package storage

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)


func DBConnection() (*pgx.Conn, error){
	err := godotenv.Load("../../.")
	if err != nil{
		log.Fatalf(err)
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatalf("Empty DB URL")
	}

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf(err)
	}
	return conn, nil
}
