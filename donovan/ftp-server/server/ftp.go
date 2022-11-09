package server

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

type FTPCustom struct {
	Address    string
	CurrentDir string
}

func New(ftp *FTPCustom) *FTPCustom {
	return &FTPCustom{
		Address:    "localhost:9001",
		CurrentDir: "/home/maxim/Projects/Go-tutorial/donovan/ftp-server",
	}
}

func (ftp *FTPCustom) Get(file string) ([]byte, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ftp *FTPCustom) Cd(dir string) (string, error) {
	if err := os.Chdir(dir); err != nil {
		return "", nil
	}
	currentDir, err := os.Getwd()
	if err != nil {
		return "", nil
	}
	return currentDir, nil
}

func (ftp *FTPCustom) Ls() ([]fs.FileInfo, error) {
	info, err := ioutil.ReadDir(ftp.CurrentDir)
	if err != nil {
		return nil, err
	}
	for _, val := range info {
		fmt.Println(val.Name())
	}
	return info, nil
}

func (ftp *FTPCustom) Close(conn interface{}) {
	c, ok := conn.(net.Conn)
	if ok {
		c.Close()
	}
}

func (ftp *FTPCustom) Run() {
	listner, err := net.Listen("tcp", ftp.Address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("FTP server start at: %s\n", ftp.Address)

	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Incoming connection: %s\n		Received data: %s\n", conn.RemoteAddr().String(), data)

		switch strings.TrimSpace(string(data)) {
		case "ls":
			res, _ := ftp.Ls()
			for _, v := range res {
				conn.Write([]byte(v.Name()))
			}
		case "cd":
			ftp.Cd("path")
		case "get":
			ftp.Get("file")
		case "close":
			ftp.Close(&conn)
		}
	}
}
