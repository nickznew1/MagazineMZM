package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/nickznew1/MagazineMZM/backend/database"
	"github.com/nickznew1/MagazineMZM/backend/internal/routes"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		fmt.Println("ошибка подключения к бд")
	}
	routes.Routes(db)
}
