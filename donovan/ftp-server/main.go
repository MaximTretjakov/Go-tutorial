package main

import (
	"fmt"
	"os"

	"github.com/MaximTretjakov/Go-tutorial/donovan/ftp-server/server"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	ftp := server.New(PORT)
	ftp.Run()
}
