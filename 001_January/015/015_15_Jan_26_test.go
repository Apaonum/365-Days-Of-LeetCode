package main

import (
	"reflect"
	"testing"
)

func TestPivotArray(t *testing.T) {
	tests := []struct {
		nums     []int
		pivot    int
		expected []int
	}{
		{
			nums:     []int{9, 12, 5, 10, 14, 3, 10},
			pivot:    10,
			expected: []int{9, 5, 3, 10, 10, 12, 14},
		},
		{
			nums:     []int{-3, 4, 3, 2},
			pivot:    2,
			expected: []int{-3, 2, 4, 3},
		},
	}

	for _, tt := range tests {
		result := pivotArray(tt.nums, tt.pivot)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("pivotArray(%v, %v) = %v, want %v", tt.nums, tt.pivot, result, tt.expected)
		}
	}
}
