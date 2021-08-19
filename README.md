# IHX Coding Challenge

## Description
This project is the solution to the IHX coding challenge provided. This repo consists of 5 different problems.

## Usage
1. Run the following command below to run the program, where `${PROBLEM_FOLDER}` is one of the problem folder we are about to run, for example: `1_swap`
    ```sh
    go run cmd/${PROBLEM_FOLDER}/main.go # go run cmd/1_swap/main.go
    ```

2. Write the input based on the problem in stdin (the input is read per line basis)

3. An output will be written in the stdout

## Problem and Solution

### 1. Swap

#### Brief Problem Description
Swap two variables without an extra variable

#### Algorithm
1. Given the variables are named `a` and `b`
2. Get the sum of the variables by adding both of variables to `a` (`a = a+b`)
3. Assign the original value `a`, by subtracting the sum (assigned to `a`) with variable `b` and assign the result to variable `b` (`b = a-b`)
4. Since `b` is now the original value of `a`, the original value of `b` can be derived by subtracting the sum with original value of `a` in `b` (`a = a-b`)

### 2. Appearance Count

#### Brief Problem Description
Given an array of integer, count the appearance of each element, print appearance of each element from the most frequent to the rarest.

#### Assumptions
The format of the input should be valid, where valid means:
- Enclosed with a square bracket (`[]`)
- The numbers should be valid as well, this means it won't contain any alphabetical character or any symbols
- Spaces are not allowed in the input

The program will consider the input as invalid if one of the above is satisfied.

Since the array may contain multiple elements that have the same appearance frequency, it's assumed that there is no specific rule on which one to be printed first to the stdout.

#### Algorithm
1. Initiate a map of integer to integer (`map[int]int{}`) that maps the element to the number of its appearance
2. Iterate the array of integers, incremene the counter by the given element inside the map
3. Use the map to create a custom struct which contains the `elem` and `count`
4. Sort the element appearance by its appearance from the most frequent to the least frequent by using built-in `sort.Slice` function provided by Go

#### Complexity
Time complexity: `O(nlogn)` due to sorting algorithm (based on quick sort)

Space complexity: `O(n)`

### 3. Merge Strings

#### Brief Problem Description
Given two different strings, merge the strings into one alternatingly (starting from the first string and followed by the second).

#### Algorithm
While we can solve this using a plain array in Golang, we can use queue data structure to help us in coordinating the characters inside the string. This also helps to clean the code for the solution

1. Generate a queue of character from 2 strings given
2. Initiate an empty string which represents the result (name this `result`)
3. While both of them is not empty pop both of them and append the first char and second char respectively to the string (`result`)
4. By this point, either of the char queues is empty (no more characters in the queue), using the non-empty queue, while it is not empty, append every characters in the queue to the resulting string.
5. The `result` will be the merged strings

#### Complexity
Time complexity: `O(m+n)`

Space complexity: `O(m+n)`

where `m` is the length of the first string and `n` is the length of the second string

### 4. Lexicographically Minimal String

#### Brief Problem Description
Given two collections of upper-case characters, merge two strings based on stack principle (only take an element at one end), so that it produces the lexicographically minimum string.

#### Algorithm
While the problem explicitly mentions the "stack" data structure, since the string has FIFO (first-in-first-out) property, queue data structure is used instead of stack.

The overall algorithm for this problem is roughly the same as the previous problem (merge strings). However, the rule for the merging of the strings is differentiated, that is by comparing the character from both of the queues. Once one of them is empty, append the rest of the non-empty queue to the resulting string.

1. Convert both of the strings into queues of characters
2. Initialize an empty string representing the result (this will be named `result`)
3. While both of the strings are not empty, compare the first element of the queue (assume character from first queue is `A` and character from second queue is `B`)
4. If `A` is "bigger" than `B`, append `B` to the `result`, pop/dequeue the element from the second queue, if `A` is "smaller" than `B`, append `A` to the `result`, and pop/dequeue the element from the first queue.
5. If `A` and `B` is an equal character, TBD
6. By this point, if either one of the queue is not empty, append the remaining characters inside the non-empty queue to `result`.
7. The `result` will be a lexicographically-minimum string.

#### Complexity
Time complexity: `O(m+n)^2` since we have to resolve conflict if `A`==`B`

Space complexity: `O(m+n)`

### 5. List Maximum
TBD

## Author
Ferdinandus Richard
