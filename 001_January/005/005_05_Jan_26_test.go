package main

import (
	"reflect"
	"testing"
)

func TestMinimumAbsDifference(t *testing.T) {
	tests := []struct {
		arr      []int
		expected [][]int
	}{
		{[]int{4, 2, 1, 3}, [][]int{{1, 2}, {2, 3}, {3, 4}}},
		{[]int{1, 3, 6, 10, 15}, [][]int{{1, 3}}},
		{[]int{3, 8, -10, 23, 19, -4, -14, 27}, [][]int{{-14, -10}, {19, 23}, {23, 27}}},
	}

	for _, tt := range tests {
		result := minimumAbsDifference(tt.arr)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("minimumAbsDifference(%v) = %v, want %v", tt.arr, result, tt.expected)
		}
	}
}
