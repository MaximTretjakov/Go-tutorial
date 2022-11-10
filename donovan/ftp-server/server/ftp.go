package server

import (
	"bufio"
	"bytes"
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

func New(port string) *FTPCustom {
	return &FTPCustom{
		Address:    port,
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

func (ftp *FTPCustom) Close() {

}

func (ftp *FTPCustom) Run() {
	listner, err := net.Listen("tcp", ftp.Address)
	if err != nil {
		log.Fatal(err)
	}
	defer listner.Close()

	fmt.Printf("FTP server start at: %s\n", ftp.Address)

	conn, err := listner.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Incoming connection: %s\n", conn.RemoteAddr().String())
		fmt.Printf("Received data: %s\n", string(data))

		switch strings.TrimSpace(string(data)) {
		case "ls":
			var b bytes.Buffer
			res, _ := ftp.Ls()
			for _, v := range res {
				b.WriteString(v.Name() + " ")
			}
			b.WriteString("\n")
			conn.Write(b.Bytes())

		case "cd":
			ftp.Cd("path")

		case "get":
			file, _ := ftp.Get("go.mod")
			f := string(file) + "\n"
			conn.Write([]byte(f))

		case "close":
			conn.Write([]byte("STOP\n"))
			return

		default:
			conn.Write([]byte("Unknown command\n"))
		}
	}
}
