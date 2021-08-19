package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwap(t *testing.T) {
	tests := []struct {
		name      string
		a         int
		b         int
		expectedA int
		expectedB int
	}{
		{
			name:      "Two positive integers",
			a:         2,
			b:         3,
			expectedA: 3,
			expectedB: 2,
		},
		{
			name:      "Positive and negative integers",
			a:         2,
			b:         -1,
			expectedA: -1,
			expectedB: 2,
		},
		{
			name:      "Two negative integers",
			a:         -200,
			b:         -1,
			expectedA: -1,
			expectedB: -200,
		},
		{
			name:      "Same integers",
			a:         5,
			b:         5,
			expectedA: 5,
			expectedB: 5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			Swap(&tc.a, &tc.b)
			assert.Equal(t, tc.a, tc.expectedA)
			assert.Equal(t, tc.b, tc.expectedB)
		})
	}
}
