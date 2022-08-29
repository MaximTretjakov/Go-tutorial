package main

import "testing"

func TestRemove(t *testing.T) {
	var tests = []struct {
		in  []int
		out []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{0},
			[]int{},
		},
		{
			[]int{1, 0, 0, 2},
			[]int{1, 2},
		},
	}

	for _, test := range tests {
		result := remove(test.in)
		for i := range result {
			if result[i] != test.out[i] {
				t.Errorf("Error. remove(%v) expect = %v, got = %v", test.in, test.out, result)
			}
		}
	}
}
