package main

import (
	"log"

	"github.com/TVolly/goapi-addresses/internal/repositories"
	"github.com/TVolly/goapi-addresses/internal/server"
)

func main() {
	cfg := server.NewConfig()
	store := repositories.NewMemoryStore()

	s := server.New(cfg, store)

	if err := s.Start(); err != nil {
		log.Fatal("Failed on start server")
	}
}
