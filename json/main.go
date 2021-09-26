package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func main() {
	fmt.Println("Work with JSON")

	author := Author{Name: "Maxim", Age: 40, Developer: true}
	book := Book{Title: "Learning GOlang", Author: author}

	fmt.Printf("%+v\n", book)

	byteArray, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(byteArray))
}
