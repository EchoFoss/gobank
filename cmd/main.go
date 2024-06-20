package main

import (
	"github.com/Fernando-Balieiro/gobank/internal/api"
	"github.com/Fernando-Balieiro/gobank/internal/infra/db"
	"log"
)

func main() {
	storage, err := db.NewPostgreDb()
	if err != nil {
		log.Fatalln(err)
	}

	server := api.NewWebServer(":8080")
	server.Start()
}
