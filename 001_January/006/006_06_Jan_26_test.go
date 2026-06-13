package main

import "testing"

func TestCandy(t *testing.T) {
	tests := []struct {
		ratings  []int
		expected int
	}{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
		{[]int{5}, 1},
		{[]int{1, 1, 1}, 3},
		{[]int{1, 2, 3}, 6},
		{[]int{3, 2, 1}, 6},
	}

	for _, tt := range tests {
		result := candy(tt.ratings)
		if result != tt.expected {
			t.Errorf("candy(%v) = %d, want %d", tt.ratings, result, tt.expected)
		}
	}
}
