package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeStrings(t *testing.T) {
	tests := []struct {
		name           string
		a              string
		b              string
		expectedResult string
	}{
		{
			name:           "Basic Test Case #1",
			a:              "saya",
			b:              "kamu",
			expectedResult: "skaaymau",
		},
		{
			name:           "Basic Test Case #2",
			a:              "kosong",
			b:              "ada",
			expectedResult: "kaodsaong",
		},
		{
			name:           "Basic Test Case #3",
			a:              "ada",
			b:              "kosong",
			expectedResult: "akdoasong",
		},
		{
			name:           "Two empty strings",
			a:              "",
			b:              "",
			expectedResult: "",
		},
		{
			name:           "One empty string",
			a:              "",
			b:              "emptystring",
			expectedResult: "emptystring",
		},
		{
			name:           "Same strings",
			a:              "helloworld",
			b:              "helloworld",
			expectedResult: "hheelllloowwoorrlldd",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mergedString := MergeStrings(tc.a, tc.b)
			assert.Equal(t, mergedString, tc.expectedResult)
		})
	}
}
