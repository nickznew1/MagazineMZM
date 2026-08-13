package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func Connect() (*pgxpool.Pool, error) {
	var db *pgxpool.Pool

	err := godotenv.Load()
	if err != nil {
		fmt.Println("No env file founded")
		return db, err
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSsl := os.Getenv("DB_SSL")

	databaseUrl := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPass, dbName, dbSsl)

	for i := 0; i < 10; i++ {
		db, err = pgxpool.New(context.Background(), databaseUrl)
		if err == nil && db.Ping(context.Background()) == nil {
			return db, nil
		}
		time.Sleep(time.Duration(i*2) * time.Second)
	}
	defer db.Close()
	if err != nil {
		log.Fatal("DB IS NOT RESPONDING")
	}
	return db, nil
}
