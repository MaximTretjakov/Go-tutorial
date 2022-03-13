package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func downloadFile(fileName string, conn net.Conn) {
	var seek int64 = 0
	const BUFFER_SIZE = 1024
	fileBuffer := make([]byte, 1024)

	file, err := os.Open(strings.TrimSpace(fileName))
	if err != nil {
		log.Fatal(err)
	}

	for {
		_, err1 := file.ReadAt(fileBuffer, seek)
		seek += BUFFER_SIZE
		fmt.Println(fileBuffer)
		conn.Write(fileBuffer)
		if err1 == io.EOF {
			break
		}
	}

	file.Close()
}

func changeDir(path string, conn net.Conn) {
	if len(path) > 0 {
		if err := os.Chdir(path); err != nil {
			log.Println(err)
		}
		cwd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		log.Println("CWD: ", cwd)
		conn.Write([]byte(cwd))
	}
}

func listing(conn net.Conn) {
	data := "ls: "
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		data += file.Name() + "\t"
	}
	conn.Write([]byte(data))
}

func handler(c net.Conn) {
	buf := make([]byte, 20)
	var command []string
	defer c.Close()

	for {
		bytes_data, err := c.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}

		data := string(buf[0:bytes_data])
		if len(data) > 0 && strings.ContainsAny(data, "_") {
			command = strings.Split(string(buf[0:bytes_data]), "_")
			if command[0] == "cd" {
				log.Printf("Remote: %s Bytes: %d Command: %s Param: %s\n", c.RemoteAddr().String(), bytes_data, command[0], command[1])
				changeDir(command[1], c)
			}
			if command[0] == "get" {
				log.Printf("Remote: %s Bytes: %d Command: %s Param: %s\n", c.RemoteAddr().String(), bytes_data, command[0], command[1])
				downloadFile(command[1], c)
			}
		} else {
			if data == "ls" {
				log.Printf("Remote: %s Bytes: %d Command: %s\n", c.RemoteAddr().String(), bytes_data, "ls")
				listing(c)
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
