package main

import (
	"fmt"
)

func romanToInt(s string) int {
	dict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0

	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && dict[s[i]] < dict[s[i+1]] {
			result -= dict[s[i]]
		} else {
			result += dict[s[i]]
		}
	}

	return result
}

func main() {
	roman1 := "III"
	roman2 := "LVIII"
	roman3 := "MCMXCIV"
	fmt.Println(romanToInt(roman1))
	fmt.Println(romanToInt(roman2))
	fmt.Println(romanToInt(roman3))
}
