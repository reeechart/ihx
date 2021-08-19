package solution

import (
	"testing"

	"github.com/reeechart/ihx/cmd/5_list_max/model"
	"github.com/stretchr/testify/assert"
)

func TestMaximumOfList(t *testing.T) {
	upperBoundOperations := make([]model.Operation, 200_000)

	for i := 0; i < 200_000; i++ {
		upperBoundOperations[i] = model.Operation{LeftBound: 1, RightBound: 1, Addition: 1_000_000_000}
	}

	tests := []struct {
		name           string
		elemCount      int
		operations     []model.Operation
		expectedResult int64
	}{
		{
			name:      "Basic Test Case",
			elemCount: 5,
			operations: []model.Operation{
				{LeftBound: 1, RightBound: 2, Addition: 100},
				{LeftBound: 2, RightBound: 5, Addition: 100},
				{LeftBound: 3, RightBound: 4, Addition: 100},
			},
			expectedResult: 200,
		},
		{
			name:      "One operation",
			elemCount: 5,
			operations: []model.Operation{
				{LeftBound: 1, RightBound: 3, Addition: 1000},
			},
			expectedResult: 1000,
		},
		{
			name:      "Equal operation",
			elemCount: 8,
			operations: []model.Operation{
				{LeftBound: 1, RightBound: 3, Addition: 1000},
				{LeftBound: 4, RightBound: 8, Addition: 1000},
			},
			expectedResult: 1000,
		},
		{
			name:           "Upper bound operations and elements",
			elemCount:      10_000_000,
			operations:     upperBoundOperations,
			expectedResult: 200_000_000_000_000,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := MaximumOfList(tc.elemCount, tc.operations)
			assert.Equal(t, result, tc.expectedResult)
		})
	}
}
