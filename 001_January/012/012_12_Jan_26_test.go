package main

import (
	"reflect"
	"testing"
)

func TestLeftRightDifference(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{10, 4, 8, 3}, []int{15, 1, 11, 22}},
		{[]int{1}, []int{0}},
		{[]int{1, 2, 3}, []int{5, 2, 3}},
		{[]int{5, 5, 5, 5}, []int{15, 5, 5, 15}},
	}

	for _, tt := range tests {
		result := leftRightDifference(tt.nums)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("leftRightDifference(%v) = %v, want %v", tt.nums, result, tt.expected)
		}
	}
}
