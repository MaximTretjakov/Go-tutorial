package main

import (
	"fmt"

	"github.com/MaximTretjakov/Go-tutorial/donovan/ex-1/tmpconv"
)

func main() {
	fmt.Println(tmpconv.Celsius(1))
	fmt.Println(tmpconv.CToK(30))
	fmt.Println(tmpconv.KToC(-243.15))
}
