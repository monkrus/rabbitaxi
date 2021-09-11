package main

import (
	"log"

	"github.com/monkrus/rabbitaxi.git/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
