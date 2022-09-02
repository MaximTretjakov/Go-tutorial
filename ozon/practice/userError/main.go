package main

import "fmt"

func handle() error {
	return &customError{}
}

type customError struct{}

func (s *customError) Error() string {
	return "Custom error"
}

func main() {
	fmt.Println(handle())
}
