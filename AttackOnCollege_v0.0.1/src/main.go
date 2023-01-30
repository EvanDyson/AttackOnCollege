package main

import (
	"CEN3031-Project/AttackOnCollege_v0.0.1/src/database"
	"CEN3031-Project/AttackOnCollege_v0.0.1/src/server"
)

func main() {
	// Initialize database
	database.Connect("database/")
	database.Migrate()

	// Initialize router
	server.StartServer()
}