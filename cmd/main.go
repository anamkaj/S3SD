package main

import (
	"direct/cmd/api"
	"direct/internal/database"
	"log"
)


func main() {

	db, err := database.PostgresConnect()
	if err != nil {
		log.Fatalln(err)
	}

	server := api.NewApiServer(":8060", db)
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}
