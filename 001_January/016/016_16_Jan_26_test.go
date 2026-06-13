package main

import "testing"

func TestMinElement(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{10, 12, 13, 14}, expected: 1},
		{nums: []int{1, 2, 3, 4}, expected: 1},
		{nums: []int{999, 19, 199}, expected: 10},
	}

	for _, tt := range tests {
		result := minElement(tt.nums)
		if result != tt.expected {
			t.Errorf("minElement(%v) = %v, want %v", tt.nums, result, tt.expected)
		}
	}
}
