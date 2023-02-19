package main

import (
	"fmt"
	"log"

	"github.com/MaximTretjakov/Go-tutorial/test-services/downloader/scraping"
	"github.com/MaximTretjakov/Go-tutorial/test-services/downloader/transformation"
)

func main() {
	rawData, err := scraping.Scraping("https://calorizator.ru/product/mushroom")
	if err != nil {
		log.Fatal(err)
	}

	data, err := transformation.Transformation(rawData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
