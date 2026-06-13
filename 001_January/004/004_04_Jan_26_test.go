package main

import "testing"

func TestMinimumOperations(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 0, 2}, {1, 0, 2}}, 0},
		{[][]int{{1, 1, 1}, {0, 0, 0}}, 3},
		{[][]int{{1}, {2}, {3}}, 2},
	}
	for _, tc := range tests {
		got := minimumOperations(tc.grid)
		if got != tc.want {
			t.Errorf("grid=%v: got %v, want %v", tc.grid, got, tc.want)
		}
	}
}
