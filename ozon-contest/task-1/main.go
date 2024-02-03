package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var count string //, data string
	_, _ = fmt.Scanln(&count)
	c, _ := strconv.Atoi(count)

	for i := 0; i < c; i++ {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			set := removeSpace(scanner.Text())
			fmt.Println(set)
			if validate(set) {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}
}

func validate(data string) bool {
	var ship1, ship2, ship3, ship4 int
	for _, v := range data {
		switch string(v) {
		case "1":
			ship1++
		case "2":
			ship2++
		case "3":
			ship3++
		case "4":
			ship4++
		}
	}
	if ship1 == 4 && ship2 == 3 && ship3 == 2 && ship4 == 1 {
		return true
	}
	return false
}

func removeSpace(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
