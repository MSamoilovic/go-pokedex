package main

import (
	"fmt"
	"testing"
)

func TesCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, _case := range cases {
		actual := cleanInput(_case.input)

		fmt.Println(actual)

		if len(actual) != len(_case.expected) {
			t.Errorf("Word counts don't match! %v vs. %v", len(actual), len(_case.expected))
		}
	}
}
