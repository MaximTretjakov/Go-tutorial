package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	var input string

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		fmt.Print("\n\nAvailable commands:\nls - list\ncd - change directory\nget - get file\nclose - close connection\n\n")
		fmt.Scanln(&input)
		if input == "close" {
			break
		}
		if len(input) > 0 {
			_, _ = conn.Write([]byte(input))
		}
	}
}
