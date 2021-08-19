package main

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/reeechart/ihx/cmd/2_appearance_count/solution"
)

var (
	inputBracketRegex = regexp.MustCompile(`^(\[.*\])$`)
)

func main() {
	inp, err := parseInput()
	if err != nil {
		panic(err)
	}

	appearance := solution.CountAndSortAppearance(inp.elems)
	solution.PrintAppearance(appearance)
}

type input struct {
	elems []int
}

// parseInput assumes the input format of [number*, number?]
func parseInput() (*input, error) {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	line = strings.Replace(line, "\n", "", -1)

	if !inputBracketRegex.MatchString(line) {
		return nil, errors.New("array format is invalid")
	}

	line = strings.Replace(line, "[", "", -1)
	line = strings.Replace(line, "]", "", -1)

	if line == "" {
		return &input{elems: []int{}}, nil
	}

	numbers := strings.Split(line, ",")

	numberArray := make([]int, len(numbers))
	for i, numstr := range numbers {
		num, err := strconv.Atoi(numstr)
		if err != nil {
			return nil, errors.New("array element is not integer-parseable")
		}

		numberArray[i] = num
	}

	return &input{elems: numberArray}, nil
}
