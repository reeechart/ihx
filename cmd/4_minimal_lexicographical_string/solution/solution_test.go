package solution

import "testing"

func TestGenerateLexicographicallyMinimalString(t *testing.T) {
	tests := []struct {
		name           string
		a              string
		b              string
		expectedResult string
	}{
		{
			name:           "Basic Test Case #1",
			a:              "JACK",
			b:              "DANIEL",
			expectedResult: "DAJACKNIEL",
		},
		{
			name:           "Basic Test Case #2",
			a:              "ABACABA",
			b:              "ABACABA",
			expectedResult: "AABABACABACABA",
		},
		{
			name:           "Empty strings",
			a:              "",
			b:              "",
			expectedResult: "",
		},
		{
			name:           "One empty string #1",
			a:              "ABACABA",
			b:              "",
			expectedResult: "ABACABA",
		},
		{
			name:           "One empty string #2",
			a:              "",
			b:              "ABACABA",
			expectedResult: "ABACABA",
		},
		{
			name:           "Smaller at one char in mid same at front with same-length strings",
			a:              "BAD",
			b:              "BCD",
			expectedResult: "BABCDD",
		},
		{
			name:           "Smaller at the end of same length string",
			a:              "HELLOWORLDA",
			b:              "HELLOWORLDB",
			expectedResult: "HEHELLLLOOWORLDAWORLDB",
		},
		{
			name:           "Substring with smaller end",
			a:              "YAY",
			b:              "YAYA",
			expectedResult: "YAYAYAY",
		},
		{
			name:           "Substring with longer end",
			a:              "BBBBABBBABC",
			b:              "BBBB",
			expectedResult: "BBBBABBBABBBBBC",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}
