package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/reeechart/ihx/pkg/queue"
)

func main() {
	inp, err := parseInput()
	if err != nil {
		panic(err)
	}

	mergedString := mergeStrings(inp.a, inp.b)
	fmt.Println(mergedString)
}

func mergeStrings(a string, b string) string {
	aq := queue.NewCharFromString(a)
	bq := queue.NewCharFromString(b)

	result := ""

	for !aq.IsEmpty() && !bq.IsEmpty() {
		charA, _ := aq.Dequeue()
		charB, _ := bq.Dequeue()

		result = fmt.Sprintf("%s%s%s", result, charA, charB)
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
