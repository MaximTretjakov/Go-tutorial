package main

import "fmt"

const base = 16777619

func Search(txt string, pattherns []string) []string {
	in := indices(txt, pattherns)
	matches := make([]string, len(in))
	i := 0
	for j, p := range pattherns {
		if _, ok := in[j]; ok {
			matches = append(matches, p)
			i++
		}
	}
	return matches
}

func indices(txt string, patterns []string) map[int]int {
	n, m := len(txt), minLen(patterns)
	matches := make(map[int]int)
	if n < m || len(patterns) == 0 {
		return matches
	}
	var mult uint32 = 1
	for i := 0; i < m-1; i++ {
		mult = (mult * base)
	}
	hp := hashPatterns(patterns, m)
	h := hash(txt[:m])
	for i := 0; i < n-m+1 && len(hp) > 0; i++ {
		if i > 0 {
			h = h - mult*uint32(txt[i-1])
			fmt.Printf("indices. first index = %d hash = %d\n", i, h)
			h = h*base + uint32(txt[i+m-1])
			fmt.Printf("indices. last index = %d hash = %d\n", i, h)
		}
		if mps, ok := hp[h]; ok {
			for _, pi := range mps {
				pat := patterns[pi]
				e := i + len(pat)
				if _, ok := matches[pi]; !ok && e <= n && pat == txt[i:e] {
					matches[pi] = i
				}
			}
		}
	}
	return matches
}

func hash(s string) uint32 {
	var h uint32
	for i := 0; i < len(s); i++ {
		h = h*base + uint32(s[i])
	}
	return h
}

func hashPatterns(patterns []string, minLen int) map[uint32][]int {
	m := make(map[uint32][]int)
	for i, t := range patterns {
		fmt.Printf("hashPatterns. slice = %v\n", t[:minLen])
		fmt.Printf("hashPatterns. hased slice = %v\n", hash(t[:minLen]))
		h := hash(t[:minLen])
		if _, ok := m[h]; ok {
			m[h] = append(m[h], i)
			fmt.Printf("exist in map m? exist = %v index = %d\n", ok, i)
		} else {
			m[h] = []int{i}
			fmt.Printf("exist in map m? exist = %v index = %d\n", ok, i)
		}
	}
	return m
}

func minLen(patterns []string) int {
	if len(patterns) == 0 {
		return 0
	}
	m := len(patterns[0])
	for i := range patterns {
		if m > len(patterns[i]) {
			m = len(patterns[i])
		}
	}
	fmt.Printf("substring min lenght. lenght = %d\n", m)
	return m
}

func main() {
	txt := "Australia is a country and continent surrounded by the Indian and Pacific oceans."
	patterns := []string{"and", "the", "surround", "Pacific", "Germany"}
	matches := Search(txt, patterns)
	for i := range matches {
		fmt.Println(matches[i])
	}
}
