package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/reeechart/ihx/cmd/4_minimal_lexicographical_string/solution"
)

func main() {
	inp, err := parseInput()
	if err != nil {
		panic(err)
	}

	for i := 0; i < inp.testCount; i++ {
		lexMinString := solution.GenerateLexicographicallyMinimalString(inp.testCases[i].a, inp.testCases[i].b)
		fmt.Println(lexMinString)
	}
}

type input struct {
	testCount int
	testCases []testCase
}

type testCase struct {
	a string
	b string
}

func parseInput() (*input, error) {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	line = strings.Replace(line, "\n", "", -1)

	t, err := strconv.Atoi(line)
	if err != nil {
		return nil, errors.New("test case count is not integer-parseable")
	}

	testCases := make([]testCase, 0)

	for i := 0; i < t; i++ {
		a, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		a = strings.Replace(a, "\n", "", -1)

		b, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		b = strings.Replace(b, "\n", "", -1)

		testCases = append(testCases, testCase{
			a: a,
			b: b,
		})
	}

	return &input{
		testCount: t,
		testCases: testCases,
	}, nil
}
