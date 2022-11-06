package main

import "github.com/MaximTretjakov/Go-tutorial/donovan/ftp-server/server"

func main() {
	ftp := server.New(&server.FTPCustom{})
	ftp.Run()
}
