package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/MaximTretjakov/Go-tutorial/test-services/downloader/entry"
)

func main() {
	url := flag.String("url", "", "Pass target url")
	flag.Parse()
	data, err := entry.Run(*url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
