package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	var db *sql.DB

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
		db, err = sql.Open("postgres", databaseUrl)
		if err == nil && db.Ping() == nil {
			return db, nil
		}
		time.Sleep(time.Duration(i*2) * time.Second)
	}
	if err != nil {
		log.Fatal("DB IS NOT RESPONDING")
	}
	return db, nil
}
