package main

import (
	"log"

	"github.com/TVolly/goapi-addresses/internal/server"
)

func main() {
	cfg := server.NewConfig()
	s := server.New(cfg)

	if err := s.Start(); err != nil {
		log.Fatal("Failed on start server")
	}
}
