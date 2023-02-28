/*
Two pointers
*/

package main

func reverseWords(s string) string {
	r := []byte(s)
	left, right := 0, 1
	for right < len(r) {
		if r[right] == 32 {
			reverse(r, left, right-1)
			left = right + 1
			right++
		}
		right++
	}
	reverse(r, left, right-1)
	return string(r)
}

func reverse(s []byte, left, right int) {
	if left >= right {
		return
	}
	s[left], s[right] = s[right], s[left]
	reverse(s, left+1, right-1)
}
