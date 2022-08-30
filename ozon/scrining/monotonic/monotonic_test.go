package main

import "testing"

func TestMonotonic(t *testing.T) {
	var tests = []struct {
		in  []int
		res bool
	}{
		{
			[]int{7, 1},
			true,
		},
		{
			[]int{23, 5, 23},
			false,
		},
	}
	for _, test := range tests {
		if res := isMonotonic(test.in); res != test.res {
			t.Errorf("isMonotonic(%v) expect = %v, got = %v", test.in, test.res, res)
		}
	}
}
