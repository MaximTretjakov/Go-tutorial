package main

import (
	"flag"

	"github.com/MaximTretjakov/Go-tutorial/donovan/ex-server/server"
)

func main() {
	address := flag.String("address", "localhost:8000", "Enter ip and port. Example: localhost:8008")
	flag.Parse()
	s := server.NewServer(address)
	s.ServerRun()
}
