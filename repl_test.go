package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "This is a test",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input:    "ALL CAPS",
			expected: []string{"all", "caps"},
		},
		{
			input:    "no caps with some number 123",
			expected: []string{"no", "caps", "with", "some", "number", "123"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Test failed, actual length does not match expected length")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Test Failed, actual word does not match expected word")
			}
		}
	}
}
