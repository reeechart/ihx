package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	inputBracketRegex = regexp.MustCompile(`^(\[.*\])$`)
)

func main() {
	inp, err := parseInput()
	if err != nil {
		panic(err)
	}

	appearance := countAppearance(inp.elems)
	printOutput(appearance)
}

func countAppearance(elems []int) []pairContainer {
	counter := make(map[int]int)
	for _, elem := range elems {
		counter[elem] += 1
	}

	appearance := make([]pairContainer, 0)
	for elem, count := range counter {
		appearance = append(appearance, pairContainer{
			elem:  elem,
			count: count,
		})
	}

	sortAppearance(appearance)

	return appearance
}

type pairContainer struct {
	elem  int
	count int
}

func sortAppearance(pairs []pairContainer) {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})
}

func printOutput(pairs []pairContainer) {
	fmt.Println("number --> total")
	for _, pair := range pairs {
		fmt.Printf("%d --> %d\n", pair.elem, pair.count)
	}
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
