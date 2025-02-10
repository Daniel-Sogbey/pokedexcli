package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "drinks like fanta are not great ",
			expected: []string{"drinks", "like", "fanta", "are", "not", "great"},
		},
		{
			input:    "  Water tanks are expensive  ",
			expected: []string{"water", "tanks", "are", "expensive"},
		},
		{
			input:    "hi world  ",
			expected: []string{"hi", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		for i := range actual {
			actualWord := actual[i]
			expectedWord := c.expected[i]

			if actualWord != expectedWord {
				t.Errorf("expected %s but got %s", expectedWord, actualWord)
			}
		}
	}
}
