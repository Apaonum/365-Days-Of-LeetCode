package main

import (
	"reflect"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	tests := []struct {
		asteroids []int
		want      []int
	}{
		{[]int{5, 10, -5}, []int{5, 10}},
		{[]int{8, -8}, []int{}},
		{[]int{10, 2, -5}, []int{10}},
		{[]int{3, 5, -6, 2, -1, 4}, []int{-6, 2, 4}},
		{[]int{-2, -1, 1, 2}, []int{-2, -1, 1, 2}},
	}
	for _, tc := range tests {
		got := asteroidCollision(tc.asteroids)
		if got == nil {
			got = []int{}
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("asteroids=%v: got %v, want %v", tc.asteroids, got, tc.want)
		}
	}
}
