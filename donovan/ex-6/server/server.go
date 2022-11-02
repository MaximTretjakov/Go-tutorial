package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type Server struct {
	ConnectionString string
	Protocol         string
}

func NewServer() *Server {
	return &Server{
		ConnectionString: "localhost:8000",
		Protocol:         "tcp",
	}
}

func (s *Server) handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func (s *Server) ServerRun() {
	listener, err := net.Listen(s.Protocol, s.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server start at %s\n", s.ConnectionString)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go s.handleConn(conn)
	}
}
