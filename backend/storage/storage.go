package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/nickznew1/MagazineMZM/backend/internal/config"

	"time"
)

func Connect(cfg *config.Config) (*pgxpool.Pool, error) {
	var db *pgxpool.Pool
	var pingDb int

	databaseUrl := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DbConfig[0].Host, cfg.DbConfig[0].Port, cfg.DbConfig[0].User,
		cfg.DbConfig[0].Password, cfg.DbConfig[0].Name, cfg.DbConfig[0].Ssl)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	db, err := pgxpool.New(ctx, databaseUrl)
	if err == nil && db.Ping(context.Background()) == nil {
		return db, nil
	}
	err = db.QueryRow(ctx, "SELECT 1").Scan(&pingDb)
	if err != nil {
		return db, err
	}
	return db, nil
}
