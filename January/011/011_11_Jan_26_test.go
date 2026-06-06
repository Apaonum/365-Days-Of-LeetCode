package main

import "testing"

func TestDistributeCandies(t *testing.T) {
	tests := []struct {
		candyType []int
		expected  int
	}{
		{[]int{1, 1, 2, 2, 3, 3}, 3},
		{[]int{1, 1, 2, 3}, 2},
		{[]int{6, 6, 6, 6}, 1},
		{[]int{1, 2}, 1},           // 2 unique types, quota 1
		{[]int{1, 2, 3, 4}, 2},     // 4 unique types, quota 2
		{[]int{-1, -1, 1, 1}, 2},   // negative values allowed, 2 unique types, quota 2
	}

	for _, tt := range tests {
		result := distributeCandies(tt.candyType)
		if result != tt.expected {
			t.Errorf("distributeCandies(%v) = %d, want %d", tt.candyType, result, tt.expected)
		}
	}
}
