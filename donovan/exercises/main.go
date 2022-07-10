package main

import (
	"bufio"
	"fmt"
)

type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
	var words int

	for {
		advance, token, err := bufio.ScanWords(p, true)
		if err != nil {
			return 0, err
		}
		if len(token) > 0 {
			words++
		}
		if advance == 0 {
			break
		}
		if advance <= len(p) {
			p = p[advance:]
		}
	}

	return words, nil
}

func main() {
	test := "Hello world Maxim\n"
	var wc WordCounter
	fmt.Println(wc.Write([]byte(test)))
}
