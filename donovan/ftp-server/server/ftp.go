package server

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net"
	"os"
)

type FTPCustom struct {
	Address    string
	CurrentDir string
}

func New(ftp *FTPCustom) *FTPCustom {
	return &FTPCustom{
		Address: "localhost:22",
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
	return info, nil
}

func (ftp *FTPCustom) Close(conn interface{}) {
	c, ok := conn.(net.Conn)
	if ok {
		c.Close()
	}
}

func (ftp *FTPCustom) Run() {
	cmd := make([]byte, 100)
	listner, err := net.Listen("tcp", ftp.Address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("FTP server start at: %s", ftp.Address)
	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		data, err := conn.Read(cmd)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Incoming connection: %s\tReceived data: %d\n", conn.RemoteAddr().String(), data)
		fmt.Printf("received command: %s\n", fmt.Sprint(data))
		switch fmt.Sprint(data) {
		case "ls":
			ftp.Ls()
		case "cd":
			ftp.Cd("path")
		case "get":
			ftp.Get("file")
		case "close":
			ftp.Close(&conn)
		}
	}
}
