package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nickznew1/MagazineMZM/backend/config"
	"github.com/nickznew1/MagazineMZM/backend/database"
	"github.com/nickznew1/MagazineMZM/backend/internal/routes"
	"log"
)

func main() {

	cfg, err = config.Load("config/config-example.yaml")
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(cfg)

	if err != nil {
		fmt.Println("ошибка подключения к бд")
	}
	routes.Routes(db, cfg)
	defer db.Close()
}
