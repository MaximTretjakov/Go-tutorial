package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	var input string
	const BUFFER_SIZE = 1024
	fromServer := make([]byte, BUFFER_SIZE)

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		fmt.Scanln(&input)

		if input == "close" {
			break
		}
		if input == "help" {
			fmt.Print("\n\nAvailable commands:\nls - list\ncd - change directory\nget - get file\nclose - close connection\n\n")
			continue
		}
		if len(input) > 0 {
			_, _ = conn.Write([]byte(input))
		}

		bytesRead, errRead := conn.Read(fromServer)
		if errRead != nil {
			log.Println(err)
		}
		fmt.Println(string(fromServer[0:bytesRead]))
	}
}
