package main

import "testing"

func TestAsteroidsDestroyed(t *testing.T) {
	tests := []struct {
		mass      int
		asteroids []int
		want      bool
	}{
		{10, []int{3, 9, 19, 5, 21}, true},
		{5, []int{4, 9, 23, 4}, false},
	}
	for _, tc := range tests {
		got := asteroidsDestroyed(tc.mass, tc.asteroids)
		if got != tc.want {
			t.Errorf("mass=%d asteroids=%v: got %v, want %v", tc.mass, tc.asteroids, got, tc.want)
		}
	}
}
