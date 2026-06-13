package main

import "testing"

func TestMinimumCost(t *testing.T) {
	tests := []struct {
		cost []int
		want int
	}{
		{[]int{1, 2, 3}, 5},
		{[]int{6, 5, 7, 9, 2, 2}, 23},
		{[]int{5, 5}, 10},
	}
	for _, tc := range tests {
		got := minimumCost(tc.cost)
		if got != tc.want {
			t.Errorf("cost=%v: got %v, want %v", tc.cost, got, tc.want)
		}
	}
}
