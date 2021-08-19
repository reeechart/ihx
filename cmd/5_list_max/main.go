package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/reeechart/ihx/cmd/5_list_max/model"
	"github.com/reeechart/ihx/cmd/5_list_max/solution"
)

func main() {
	inp, err := parseInput()
	if err != nil {
		panic(err)
	}

	listMax := solution.MaximumOfList(inp.elemCount, inp.operations)
	fmt.Println(listMax)
}

type input struct {
	elemCount      int
	operationCount int
	operations     []model.Operation
}

func parseInput() (*input, error) {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	line = strings.Replace(line, "\n", "", -1)

	numbers := strings.Split(line, " ")

	if len(numbers) != 2 {
		return nil, errors.New("N and M value is invalid")
	}

	elemCount, err := strconv.Atoi(numbers[0])
	if err != nil {
		return nil, errors.New("element count (N) is not integer-parseable")
	}

	operationCount, err := strconv.Atoi(numbers[1])
	if err != nil {
		return nil, errors.New("operation count (M) is not integer-parseable")
	}

	operations := make([]model.Operation, operationCount)
	for i := 0; i < operationCount; i++ {
		operation, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		operation = strings.Replace(operation, "\n", "", -1)

		ops := strings.Split(operation, " ")
		if len(ops) != 3 {
			return nil, errors.New("operation format is invalid")
		}

		a, err := strconv.Atoi(ops[0])
		if err != nil {
			return nil, errors.New("value of A is not integer-parseable")
		}

		b, err := strconv.Atoi(ops[1])
		if err != nil {
			return nil, errors.New("value of B is not integer-parseable")
		}

		k, err := strconv.ParseInt(ops[2], 10, 64)
		if err != nil {
			return nil, err
		}

		operations[i] = model.Operation{
			LeftBound:  a,
			RightBound: b,
			Addition:   k,
		}
	}

	return &input{
		elemCount:      elemCount,
		operationCount: operationCount,
		operations:     operations,
	}, nil
}
