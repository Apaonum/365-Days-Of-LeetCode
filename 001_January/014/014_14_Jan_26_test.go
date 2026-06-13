package main

import (
	"reflect"
	"testing"
)

func TestGetResults(t *testing.T) {
	tests := []struct {
		queries  [][]int
		expected []bool
	}{
		{
			queries:  [][]int{{1, 2}, {2, 3, 3}, {2, 3, 1}, {2, 2, 2}},
			expected: []bool{false, true, true},
		},
		{
			queries:  [][]int{{1, 7}, {2, 7, 6}, {1, 2}, {2, 7, 5}, {2, 7, 6}},
			expected: []bool{true, true, false},
		},
	}

	for _, tt := range tests {
		result := getResults(tt.queries)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("getResults(%v) = %v, want %v", tt.queries, result, tt.expected)
		}
	}
}
