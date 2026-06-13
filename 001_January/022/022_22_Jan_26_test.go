package main

import "testing"

func TestWeightedWordMapping(t *testing.T) {
	tests := []struct {
		words    []string
		weights  []int
		expected string
	}{
		// Example 1
		{
			words:    []string{"abcd", "def", "xyz"},
			weights:  []int{5, 3, 12, 14, 1, 2, 3, 2, 10, 6, 6, 9, 7, 8, 7, 10, 8, 9, 6, 9, 9, 8, 3, 7, 7, 2},
			expected: "rij",
		},
		// Example 2: all weight 1, each single-char word maps to mod 1 -> 'y'
		{
			words:    []string{"a", "b", "c"},
			weights:  []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			expected: "yyy",
		},
		// Example 3
		{
			words:    []string{"abcd"},
			weights:  []int{7, 5, 3, 4, 3, 5, 4, 9, 4, 2, 2, 7, 10, 2, 5, 10, 6, 1, 2, 2, 4, 1, 3, 4, 4, 5},
			expected: "g",
		},
		// mod 0 -> maps to 'z'
		{
			words:    []string{"a"},
			weights:  []int{26, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			expected: "z",
		},
	}

	for _, tt := range tests {
		result := weightedWordMapping(tt.words, tt.weights)
		if result != tt.expected {
			t.Errorf("weightedWordMapping(%v, ...) = %q, want %q", tt.words, result, tt.expected)
		}
	}
}
