package main

import (
	"log"

	"github.com/TVolly/goapi-addresses/internal/repositories"
	"github.com/TVolly/goapi-addresses/internal/server"
)

func main() {
	dbURL := "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable"
	cfg := server.NewConfig()
	store := repositories.NewSqlStore(dbURL)

	s := server.New(cfg, store)

	if err := s.Start(); err != nil {
		log.Fatal("Failed on start server")
	}
}
