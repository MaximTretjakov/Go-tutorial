package main

import "testing"

func TestLenghtOfLastWord(t *testing.T) {
	var tests = []struct {
		testString string
		lenght     int
	}{
		{
			"Hello World",
			5,
		},
		{
			"   fly me   to   the moon  ",
			4,
		},
		{
			"luffy is still joyboy",
			6,
		},
		{
			"World",
			5,
		},
	}
	for _, test := range tests {
		if wordLenght := lengthOfLastWord(test.testString); wordLenght != test.lenght {
			t.Errorf("lengthOfLastWord(%s)=%d", test.testString, test.lenght)
		}
	}
}
