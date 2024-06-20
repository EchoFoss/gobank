package main

import (
	"github.com/Fernando-Balieiro/gobank/internal/api"
)

func main() {

	server := api.NewWebServer(":8080")
	server.Start()
}
