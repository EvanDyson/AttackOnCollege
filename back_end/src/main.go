package main

import (
	"AttackOnCollege/back_end/src/database"
	"AttackOnCollege/back_end/src/server"
)

//"AttackOnCollege/back_end/src/database"
//"AttackOnCollege/back_end/src/server"

func main() {
	// Initialize database
	database.Connect("./back_end/src/database/")

	// DATABASE PATH FOR DEBUGGING
	// database.Connect("./database/")

	database.Migrate()

	// Initialize router
	server.StartServer()

}
