package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/reeechart/ihx/cmd/3_string_merge/solution"
)

func main() {
	inp, err := parseInput()
	if err != nil {
		panic(err)
	}

	mergedString := solution.MergeStrings(inp.a, inp.b)
	fmt.Println(mergedString)
}

type input struct {
	a string
	b string
}

func parseInput() (*input, error) {
	reader := bufio.NewReader(os.Stdin)
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

	return &input{a: a, b: b}, nil
}
