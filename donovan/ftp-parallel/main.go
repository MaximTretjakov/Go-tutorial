package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func changeDir(path string) {
	if len(path) > 0 {
		if err := os.Chdir(path); err != nil {
			log.Println(err)
		}
		cwd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		log.Println("CWD: ", cwd)
	}
}

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
	var commandAndData []string
	defer c.Close()

	for {
		bytes_data, err := c.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}

		data := string(buf[0:bytes_data])
		if len(data) > 0 && strings.ContainsAny(data, "_") {
			commandAndData = strings.Split(string(buf[0:bytes_data]), "_")
			if commandAndData[0] == "cd" {
				log.Printf("Remote: %s Bytes: %d Command: %s Param: %s\n", c.RemoteAddr().String(), bytes_data, commandAndData[0], commandAndData[1])
				changeDir(commandAndData[1])
			}
			if commandAndData[0] == "get" {
				log.Printf("Remote: %s Bytes: %d Command: %s Param: %s\n", c.RemoteAddr().String(), bytes_data, commandAndData[0], commandAndData[1])
			}
		} else {
			if data == "ls" {
				log.Printf("Remote: %s Bytes: %d Command: %s\n", c.RemoteAddr().String(), bytes_data, "ls")
				listing()
			}
			if data == "close" {
				log.Printf("Remote: %s Bytes: %d Command: %s\n", c.RemoteAddr().String(), bytes_data, "close")
			}
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
