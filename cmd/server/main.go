package main

import (
	"log"

	"github.com/JQweenq/go-back/internal/app/server"
)

func main() {
	config := server.NewConfig()
	s := server.New()

	if err := s.Start(config); err != nil {
		log.Fatal(err)
	}
}
