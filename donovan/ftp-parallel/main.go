package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

// func changeDir(path string) {
// 	if len(path) > 0 {
// 		if err := os.Chdir(path); err != nil {
// 			log.Println(err)
// 		}
// 	}
// }

func listing() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func handler(c net.Conn) {
	buf := make([]byte, 20)
	defer c.Close()

	for {
		bytes_data, err := c.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}

		switch string(buf[0:bytes_data]) {
		case "ls":
			log.Printf("Remote: %s Bytes: %d Command: %s\n", c.RemoteAddr().String(), bytes_data, "ls")
			listing()
		case "cd":
			log.Printf("Remote: %s Bytes: %d Command: %s\n", c.RemoteAddr().String(), bytes_data, "cd")
			// changeDir()
		case "get":
			log.Printf("Remote: %s Bytes: %d Command: %s\n", c.RemoteAddr().String(), bytes_data, "get")
		case "close":
			log.Printf("Remote: %s Bytes: %d Command: %s\n", c.RemoteAddr().String(), bytes_data, "close")
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handler(conn)
	}
}
