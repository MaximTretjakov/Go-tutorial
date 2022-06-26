package main

import "testing"

func TestSummOk(t *testing.T) {
	s := Summ(2, 2)
	if s != 4 {
		t.Errorf(`2+2=4 got %v`, s)
	}
}

func TestSummMoreCases(t *testing.T) {
	var tests = []struct {
		x    int
		y    int
		summ int
	}{
		{
			2,
			2,
			4,
		},
		{
			2,
			2,
			5,
		},
	}

	for _, test := range tests {
		s := Summ(test.x, test.y)
		if s != test.summ {
			t.Errorf(`%v + %v = %v got %v`, test.x, test.y, test.summ, s)
		}
	}
}
