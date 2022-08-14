package main

import (
	"fmt"
	"log"

	"github.com/JQweenq/go-back/internal/app/server"
)

func main() {
	config := server.NewConfig()

	fmt.Printf("%+v\n", config)

	s := server.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
