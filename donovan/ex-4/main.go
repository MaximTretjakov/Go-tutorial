package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type WordCounter int
type StringCounter int
type CountinWriter struct{}

func (cw *CountinWriter) Write(p []byte) (int, error) { return len(p), nil }

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
	var buf bytes.Buffer
	_, err := fmt.Fprintln(&buf)
	if err != nil {
		fmt.Println(err)
	}
	return &buf, nil
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
	var cw CountinWriter
	cw.Write([]byte(s))
	// w, counterBytes := CountingWriter(&cw)
	// fmt.Printf("New Writer\t%v\tBytes counter\t%d\n", w, *counterBytes)
	fmt.Println(CountingWriter(&cw))
}
