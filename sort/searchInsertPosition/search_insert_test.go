package main

import "testing"

func TestSearchInsert(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		index  int
	}{
		{
			[]int{1, 3, 5, 6},
			5,
			2,
		},
		{
			[]int{1, 3, 5, 6},
			2,
			1,
		},
		{
			[]int{1, 3, 5, 6},
			7,
			4,
		},
	}

	for _, test := range tests {
		if idx := searchInsert(test.nums, test.target); idx != test.index {
			t.Errorf("searchInsert(%v) = %d", test.nums, test.target)
		}
	}
}
