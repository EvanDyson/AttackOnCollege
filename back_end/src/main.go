package main

import (
	"CEN3031-Project/back_end/src/database"
	"CEN3031-Project/back_end/src/server"
)

func main() {
	// Initialize database
	database.Connect("./back_end/src/database/")
	database.Migrate()

	// Initialize router
	server.StartServer()
}
