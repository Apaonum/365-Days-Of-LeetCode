package main

import "testing"

func TestEarliestFinishTime(t *testing.T) {
	tests := []struct {
		landStartTime  []int
		landDuration   []int
		waterStartTime []int
		waterDuration  []int
		expected       int
	}{
		{[]int{2, 8}, []int{4, 1}, []int{6}, []int{3}, 9},
		{[]int{5}, []int{3}, []int{1}, []int{10}, 14},
		{[]int{1}, []int{1}, []int{1}, []int{1}, 3},  // both open at 1, land then water: 1+1=2, 2+1=3
		{[]int{10}, []int{5}, []int{1}, []int{2}, 15}, // water first: 1+2=3, wait, land at 10+5=15
	}

	for _, tt := range tests {
		result := earliestFinishTime(tt.landStartTime, tt.landDuration, tt.waterStartTime, tt.waterDuration)
		if result != tt.expected {
			t.Errorf("earliestFinishTime(%v, %v, %v, %v) = %d, want %d",
				tt.landStartTime, tt.landDuration, tt.waterStartTime, tt.waterDuration, result, tt.expected)
		}
	}
}
