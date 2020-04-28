package main

import (
	"log"

	"github.com/TVolly/goapi-addresses/internal/repositories"
	"github.com/TVolly/goapi-addresses/internal/server"
)

func main() {
	cfg := server.NewConfig()
	repoRegistry := repositories.NewMemoryRegistry()

	s := server.New(cfg, repoRegistry)

	if err := s.Start(); err != nil {
		log.Fatal("Failed on start server")
	}
}
