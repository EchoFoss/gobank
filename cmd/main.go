package main

import (
	"flag"
	"fmt"
	"github.com/Fernando-Balieiro/gobank/internal/api"
	"github.com/Fernando-Balieiro/gobank/internal/infra/db"
	scripts "github.com/Fernando-Balieiro/gobank/scripts/db"
	"log"
)

func main() {
	seeding := flag.Bool("seed", false, "seed the db")
	flag.Parse()
	storage, err := db.NewPostgreDb()
	if err != nil {
		log.Fatalln(err)
	}

	if err := storage.Init(); err != nil {
		log.Fatalln(err)
	}

	if *seeding {
		fmt.Println("seeding the database...")
		scripts.SeedAccounts(storage)
	}

	api.NewWebServer(":8080", storage).Start()
}
