package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAppearance(t *testing.T) {
	tests := []struct {
		name          string
		arr           []int
		expectedCount []PairContainer
	}{
		{
			name: "Successfully counting multiple different elements",
			arr:  []int{4, 6, 3, 5, 4, 6, 7, 8, 3, 4, 6, 7, 5, 4, 6, 4, 4, 5, 6},
			expectedCount: []PairContainer{
				{Elem: 4, Count: 6},
				{Elem: 6, Count: 5},
				{Elem: 5, Count: 3},
				{Elem: 3, Count: 2},
				{Elem: 7, Count: 2},
				{Elem: 8, Count: 1},
			},
		},
		{
			name:          "Successfully counting multiple same elements",
			arr:           []int{1, 1, 1, 1, 1, 1, 1, 1},
			expectedCount: []PairContainer{{Elem: 1, Count: 8}},
		},
		{
			name:          "Successfully counting a single element",
			arr:           []int{4},
			expectedCount: []PairContainer{{Elem: 4, Count: 1}},
		},
		{
			name:          "Successfully counting zero element",
			arr:           []int{},
			expectedCount: []PairContainer{},
		},
		{
			name: "Successfully a single different elements",
			arr:  []int{1, 2, 3, 4, 5, 6, 7},
			expectedCount: []PairContainer{
				{Elem: 1, Count: 1},
				{Elem: 2, Count: 1},
				{Elem: 3, Count: 1},
				{Elem: 4, Count: 1},
				{Elem: 5, Count: 1},
				{Elem: 6, Count: 1},
				{Elem: 7, Count: 1},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			count := countAppearance(tc.arr)
			if !assert.ElementsMatch(t, count, tc.expectedCount) {
				t.Errorf("counter mismatch, got: %v, want: %v", count, tc.expectedCount)
			}
		})
	}
}
