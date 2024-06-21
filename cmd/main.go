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

	if err := storage.Init(); err != nil {
		log.Fatalln(err)
	}

	api.NewWebServer(":8080", storage).Start()
}
