package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	type testCase struct {
		input    string
		expected []string
	}

	cases := []testCase{
		{
			input:    "   hello world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// check resultant slice first
		if len(actual) != len(c.expected) {
			t.Errorf("len doesn't match - actual len: %v, expected len: %v", len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			// check if words match
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v - expected %v", c.input, actual, c.expected)
			}
		}
	}
}
