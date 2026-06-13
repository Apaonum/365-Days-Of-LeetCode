package main

import "testing"

func TestMaxSubarrayValue(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int64
	}{
		{nums: []int{1, 3, 2}, k: 2, expected: 4},
		{nums: []int{4, 2, 5, 1}, k: 3, expected: 12},
	}

	for _, tt := range tests {
		result := maxSubarrayValue(tt.nums, tt.k)
		if result != tt.expected {
			t.Errorf("maxSubarrayValue(%v, %v) = %v, want %v", tt.nums, tt.k, result, tt.expected)
		}
	}
}
