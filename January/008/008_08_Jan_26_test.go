package main

import "testing"

func TestEarliestFinishTime(t *testing.T) {
	tests := []struct {
		name          string
		landStartTime []int
		landDuration  []int
		waterStartTime []int
		waterDuration  []int
		want          int
	}{
		{
			name:           "example 1 - land then water optimal",
			landStartTime:  []int{2, 8},
			landDuration:   []int{4, 1},
			waterStartTime: []int{6},
			waterDuration:  []int{3},
			want:           9,
		},
		{
			name:           "example 2 - water then land optimal",
			landStartTime:  []int{5},
			landDuration:   []int{3},
			waterStartTime: []int{1},
			waterDuration:  []int{10},
			want:           14,
		},
		{
			name:           "single ride each, land opens first",
			landStartTime:  []int{1},
			landDuration:   []int{2},
			waterStartTime: []int{5},
			waterDuration:  []int{1},
			want:           6,
		},
		{
			name:           "rides open at same time",
			landStartTime:  []int{3},
			landDuration:   []int{2},
			waterStartTime: []int{3},
			waterDuration:  []int{2},
			want:           7,
		},
		{
			name:           "water ride finishes exactly when land opens",
			landStartTime:  []int{10},
			landDuration:   []int{5},
			waterStartTime: []int{1},
			waterDuration:  []int{9},
			want:           15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := earliestFinishTime(tt.landStartTime, tt.landDuration, tt.waterStartTime, tt.waterDuration)
			if got != tt.want {
				t.Errorf("earliestFinishTime() = %d, want %d", got, tt.want)
			}
		})
	}
}
