package main

import (
	"testing"
)

func test_add_two(t *testing.T) {
	var tests = []struct {
		input  int
		output int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{5, 3},
	}

	for _, test := range tests {
		if output := add_two(test.input); output != test.output {
			t.Error("Test failed")
		}
	}
}
