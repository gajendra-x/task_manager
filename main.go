package main

import (
	"task_manager/config/db"
	"task_manager/seed"
)

func main() {

	db.InitializeDatabase()
	seed.SeedUsers()
	// app.Server(constants.PORT)
}
