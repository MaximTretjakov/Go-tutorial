package main

import "testing"

func TestPlusOne(t *testing.T) {
	var tests = []struct {
		input  []int
		output []int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 4},
		},
		{
			[]int{3},
			[]int{4},
		},
		{
			[]int{9, 9},
			[]int{1, 0, 0},
		},
	}
	for _, test := range tests {
		result := plusOne(test.input)
		if len(result) != len(test.output) {
			t.Errorf("slice len expect = %d, got = %d", len(test.output), len(result))
		}
		for i := range test.output {
			if test.output[i] != result[i] {
				t.Errorf("slice not equal. expect = %v, got = %v", test.output, result)
			}
		}
	}
}
