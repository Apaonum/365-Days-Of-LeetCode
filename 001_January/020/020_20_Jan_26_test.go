package main

import "testing"

func TestNumberWays(t *testing.T) {
	tests := []struct {
		edges    [][]int
		expected int
	}{
		// Single edge — only 1 valid assignment (weight 1)
		{edges: [][]int{{1, 2}}, expected: 1},
		// Max depth is 2; 2^(2-1) = 2
		{edges: [][]int{{1, 2}, {1, 3}, {3, 4}, {3, 5}}, expected: 2},
		// Linear chain depth 3; 2^(3-1) = 4
		{edges: [][]int{{1, 2}, {2, 3}, {3, 4}}, expected: 4},
		// Star graph — max depth is 1; 2^(1-1) = 1
		{edges: [][]int{{1, 2}, {1, 3}, {1, 4}}, expected: 1},
	}

	for _, tt := range tests {
		result := numberWays(tt.edges)
		if result != tt.expected {
			t.Errorf("numberWays(%v) = %v, want %v", tt.edges, result, tt.expected)
		}
	}
}
