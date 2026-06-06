package main

import "testing"

func TestTotalWaviness(t *testing.T) {
	tests := []struct {
		num1, num2 int
		expected   int
	}{
		{120, 130, 3},
		{198, 202, 3},
		{4848, 4848, 2},
		{1, 99, 0},   // all 1-2 digit numbers have waviness 0
		{100, 100, 0}, // 1-0-0: middle digit 0 is not peak (0 < 1 but 0 == 0), not valley
	}

	for _, tt := range tests {
		result := totalWaviness(tt.num1, tt.num2)
		if result != tt.expected {
			t.Errorf("totalWaviness(%d, %d) = %d, want %d", tt.num1, tt.num2, result, tt.expected)
		}
	}
}
