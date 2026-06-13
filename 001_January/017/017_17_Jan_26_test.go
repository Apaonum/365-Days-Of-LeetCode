package main

import (
	"reflect"
	"testing"
)

func TestStringIndices(t *testing.T) {
	tests := []struct {
		wordsContainer []string
		wordsQuery     []string
		expected       []int
	}{
		{
			wordsContainer: []string{"abcd", "bcd", "xbcd"},
			wordsQuery:     []string{"cd", "bcd", "xyz"},
			expected:       []int{1, 1, 1},
		},
		{
			wordsContainer: []string{"abcdefgh", "poiuygh", "ghghgh"},
			wordsQuery:     []string{"gh", "acbfgh", "acbfegh"},
			expected:       []int{2, 0, 2},
		},
	}

	for _, tt := range tests {
		result := stringIndices(tt.wordsContainer, tt.wordsQuery)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("stringIndices(%v, %v) = %v, want %v", tt.wordsContainer, tt.wordsQuery, result, tt.expected)
		}
	}
}
