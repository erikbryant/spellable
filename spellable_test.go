package main

import (
	"testing"
)

func TestSpellable(t *testing.T) {
	testCases := []struct {
		w1       string
		w2       string
		expected bool
	}{
		{"catch", "cat", true},
		{"think", "ink", true},
		{"think", "think", true},
		{"think", "tthink", false},
		{"think", "thinkk", false},
		{"think", "kniht", true},
		{"", "tthink", false},
		{"think", "", true},
	}

	for _, testCase := range testCases {
		answer := spellable(testCase.w1, testCase.w2)
		if answer != testCase.expected {
			t.Errorf("For %v, %v expected %v, got %v", testCase.w1, testCase.w2, testCase.expected, answer)
		}
	}
}
