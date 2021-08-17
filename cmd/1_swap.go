package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inp, err := parseInput()
	if err != nil {
		panic(err)
	}

	a := inp.a
	b := inp.b

	fmt.Printf("Before swap: a=%d, b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap: a=%d, b=%d\n", a, b)
}

// Swap function swaps two integers in-place.
func swap(a *int, b *int) {
	*a += *b
	*b = *a - *b
	*a = *a - *b
}

type input struct {
	a int
	b int
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
		return nil, errors.New("input length must be 2")
	}

	fmt.Println(numbers)

	a, err := strconv.Atoi(numbers[0])
	if err != nil {
		return nil, errors.New("first number is not integer-parseable")
	}

	b, err := strconv.Atoi(numbers[1])
	if err != nil {
		return nil, errors.New("second number is not integer-parseable")
	}

	return &input{a, b}, nil
}
