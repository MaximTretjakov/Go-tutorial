package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type Server struct {
	IP   string
	Port string
}

func NewServer() *Server {
	return &Server{
		IP:   "localhost",
		Port: "8000",
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
	listener, err := net.Listen(s.IP, s.Port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server start. Port: %s IP: %s", s.Port, s.IP)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go s.handleConn(conn)
	}
}
