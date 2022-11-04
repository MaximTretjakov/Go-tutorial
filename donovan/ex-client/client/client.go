package client

import (
	"io"
	"log"
	"net"
	"os"
)

func MoveCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func Client(address *string) {
	conn, err := net.Dial("tcp", *address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	MoveCopy(os.Stdout, conn)
}
