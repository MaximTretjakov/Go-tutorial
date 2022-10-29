package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type WordCounter int
type StringCounter int

func (wc WordCounter) Write(p []byte) (int, error) {
	input := string(p)
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanWords)
	counter := 0
	for s.Scan() {
		counter++
	}
	return counter, nil
}

func (ws *StringCounter) Write(p []byte) int {
	input := string(p)
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanLines)
	counter := 0
	for s.Scan() {
		counter++
	}
	return counter
}

// todo
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var new io.Writer
	b, err := fmt.Fprintln(w, new)
	if err != nil {
		fmt.Println(err)
	}
	nunBytes := int64(b)
	return new, &nunBytes
}

func main() {
	s := "Hello, мир!"
	str := "Test\n Test test Hello ! \n test"
	var b WordCounter
	var d StringCounter

	// StringCounter
	fmt.Printf("Strings\t%d\n", d.Write([]byte(str)))

	// WordCounter
	r, err := b.Write([]byte(s))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Words\t%d\n", r)

	// CountingWriter
	w, counterBytes := CountingWriter(b)
	fmt.Printf("New Writer\t%v\tBytes counter\t%d\n", w, *counterBytes)
}
