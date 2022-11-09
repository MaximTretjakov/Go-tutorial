package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9001")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(">>>")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->: " + message)
	}

}
