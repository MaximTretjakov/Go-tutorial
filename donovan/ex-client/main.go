package main

import (
	"flag"

	"github.com/MaximTretjakov/Go-tutorial/donovan/ex-client/client"
)

func main() {
	address := flag.String("address", "localhost:8000", "Enter server address/port")
	flag.Parse()
	client.Client(address)
}
