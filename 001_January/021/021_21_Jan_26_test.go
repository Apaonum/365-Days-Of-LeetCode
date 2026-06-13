package main

import (
	"reflect"
	"testing"
)

func TestAnswerQueries(t *testing.T) {
	tests := []struct {
		edges    [][]int
		queries  [][]int
		expected []int
	}{
		// Example 1: single edge
		{
			edges:    [][]int{{1, 2}},
			queries:  [][]int{{1, 1}, {1, 2}},
			expected: []int{0, 1},
		},
		// Example 2: tree with depth 2
		{
			edges:    [][]int{{1, 2}, {1, 3}, {3, 4}, {3, 5}},
			queries:  [][]int{{1, 4}, {3, 4}, {2, 5}},
			expected: []int{2, 1, 4},
		},
		// Same node query — path length 0, always 0 valid assignments
		{
			edges:    [][]int{{1, 2}, {2, 3}},
			queries:  [][]int{{2, 2}},
			expected: []int{0},
		},
		// Path of length 1 — exactly 1 valid assignment
		{
			edges:    [][]int{{1, 2}, {2, 3}},
			queries:  [][]int{{2, 3}},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		result := answerQueries(tt.edges, tt.queries)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("answerQueries(%v, %v) = %v, want %v", tt.edges, tt.queries, result, tt.expected)
		}
	}
}
