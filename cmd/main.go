package main

import (
	"database/sql"
	"go-ecommerce/cmd/api"
	"go-ecommerce/db"
	"log"
)

func main() {

	db, err := db.NewSQLiteStorage("./data.db")
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database!")
}
