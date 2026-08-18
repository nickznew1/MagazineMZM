package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/nickznew1/MagazineMZM/config"
	"gopkg.in/yaml.v3"
	"log"
	"os"

	"time"
)

func Connect(cfg *config.Config) (*pgxpool.Pool, error) {
	var db *pgxpool.Pool

	/*var cfg Config

	data, err := os.ReadFile("config/config-example.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatal(err)
	}*/

	databaseUrl := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DbConfig[0].Host, cfg.DbConfig[0].Port, cfg.DbConfig[0].User, cfg.DbConfig[0].Password, cfg.DbConfig[0].Name, cfg.DbConfig[0].Ssl)

	fmt.Println(databaseUrl)

	for i := 0; i < 10; i++ {
		db, err = pgxpool.New(context.Background(), databaseUrl)
		if err == nil && db.Ping(context.Background()) == nil {
			fmt.Println(err)
			return db, nil
		}
		time.Sleep(time.Duration(i*2) * time.Second)
	}
	if err != nil {
		log.Fatal("DB IS NOT RESPONDING")
	}
	return db, nil
}
