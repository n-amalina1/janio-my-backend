package main

import (
	"database/sql"

	routes "janio-my-backend/routes"
)

var db *sql.DB

func main() {
	routes.SetupRoutes(db)
}
