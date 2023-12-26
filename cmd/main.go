package main

import (
	"log"
	"users/internal/api"
	"users/internal/config"
	"users/internal/data/memoryStore"
)

func main() {
	store := memoryStore.NewMemoryStore()
	server := api.NewServer(store)

	port := config.Config("PORT")
	err := server.Start(port)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
