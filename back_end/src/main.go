package main

import (
	"CEN3031-Project/back_end/src/database"
	"CEN3031-Project/back_end/src/server"
)

func main() {
	// Initialize database
	// Change paths once we start running the server with npm
	database.Connect("./back_end/src/database/")
	database.Migrate()

	// Initialize router
	server.StartServer()
}
