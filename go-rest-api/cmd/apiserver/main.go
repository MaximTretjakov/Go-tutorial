package main

import (
	"flag"
	"log"

	"github.com/MaximTretjakov/Go-tutorial/go-rest-api/internal/app/apiserver"
)

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
