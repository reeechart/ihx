package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/reeechart/ihx/pkg/queue"
)

func main() {
	inp, err := parseInput()
	if err != nil {
		panic(err)
	}

	for i := 0; i < inp.testCount; i++ {
		lexMinString := generateLexicographicallyMinimalString(inp.testCases[i].a, inp.testCases[i].b)
		fmt.Println(lexMinString)
	}
}

func generateLexicographicallyMinimalString(a string, b string) string {
	aq := queue.NewCharFromString(a)
	bq := queue.NewCharFromString(b)

	result := ""

	for !aq.IsEmpty() && !bq.IsEmpty() {
		charA, _ := aq.Peek()
		charB, _ := bq.Peek()

		switch strings.Compare(charA, charB) {
		case 1:
			result = fmt.Sprintf("%s%s", result, charB)
			bq.Dequeue()
		case -1:
			result = fmt.Sprintf("%s%s", result, charA)
			aq.Dequeue()
		default: // case 0
			// TODO: Decide which side to pick by peeking the next element, if all matches, pick any side
			result = fmt.Sprintf("%s%s", result, charA)
			aq.Dequeue()
		}
	}

	for !aq.IsEmpty() {
		char, _ := aq.Dequeue()
		result = fmt.Sprintf("%s%s", result, char)
	}

	for !bq.IsEmpty() {
		char, _ := bq.Dequeue()
		result = fmt.Sprintf("%s%s", result, char)
	}

	return result
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
