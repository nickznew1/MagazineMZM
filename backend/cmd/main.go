package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"myWebApp/backend/database"
	"myWebApp/backend/internal/routes"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		fmt.Println("ошибка подключения к бд")
	}
	routes.Routes(db)
}
