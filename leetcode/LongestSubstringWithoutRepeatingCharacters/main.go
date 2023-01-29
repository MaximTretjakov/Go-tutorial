/*
Задача lengthOfLongestSubstring решается паттерном sliding window
*/
package main

import "fmt"

// func lengthOfLongestSubstring(s string) int {
// 	startWindow := 0
// 	// endWindow := 0
// 	longestSubstringLenght := 0
// 	currentSubstringLenght := 0
// 	store := make(map[byte]int, len(s))

// 	for i := 0; i < len(s)-1; i++ {
// 		// endWindow = i
// 		if val, ok := store[s[i]]; ok {
// 			if val >= startWindow {
// 				currentSubstringLenght++
// 			} else {
// 				store[s[i]] = i
// 				startWindow++
// 			}
// 			if currentSubstringLenght > longestSubstringLenght {
// 				longestSubstringLenght = currentSubstringLenght
// 			}
// 			startWindow++
// 		} else {
// 			store[s[i]] = i
// 		}
// 	}
// 	return longestSubstringLenght
// }

func lengthOfLongestSubstring(s string) int {
    result := 0
    // мапа для отслеживания повторяющихся элементов
    seen := map[byte]bool{}
    
    // левая граница окна
    start := 0
    
    for end := 0; end < len(s); end++ {
        // если элемент есть в мапе то делаем false для элемента с индексом start
        // увеличиваем start сдвигаем левый край окна
        for seen[s[end]] {
            seen[s[start]] = false
            start++
        }
        
        // если элемента нет добавляем
        seen[s[end]] = true;
        
        // тут подсчет элементов
        if end - start + 1 > result {
            result = end - start + 1
        }
    }
    
    return result
}

func main() {
	// s1 := "abcabcbb" // valid output = 3
	// fmt.Println(lengthOfLongestSubstring(s1))
	// s2 := "bbbbb" // valid output = 1
	// fmt.Println(lengthOfLongestSubstring(s2))
	s3 := "pwwkew" // valid output = 3
	fmt.Println(lengthOfLongestSubstring(s3))
}
