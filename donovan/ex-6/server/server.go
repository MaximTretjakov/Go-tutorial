package server

import (
	"io"
	"log"
	"net"
	"time"
)

type Server struct{
	IP string
	Port string
}

func(s *Server) handleConn(c net.Conn){
	defer c.Close()
	for{
		_, err:=io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil{
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func (s *Server) Server() {
	listener, err := net.Listen("localhost", "8000")
	if err != nil {
		log.Fatal(err)
	}

	for{
		conn, err:=listener.Accept()
		if err!= nil{
			log.Print(err)
			continue
		}
		go s.handleConn(conn)
	}
}
