package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, msg)
	}
}

func handleConn(c net.Conn) {
	ch := make(chan string)
	go clientWriter(c, ch)

	who := c.RemoteAddr().String()
	ch <- "Вы: " + who + "\n"
	messages <- who + " подключился\n"
	entering <- ch

	input := bufio.NewScanner(c)
	for input.Scan() {
		messages <- who + ": " + input.Text() + "\n"
	}

	leaving <- ch
	messages <- who + " отключился\n"
	c.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
