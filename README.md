# IHX Coding Challenge

## Description
This project is the solution to the IHX coding challenge provided. This repo consists of 5 different problems.

## Usage
```sh
go run cmd/${PROBLEM_FOLDER}/main.go
```

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

### 3. Merge Strings
TBD

### 4. Lexicographically Minimal String
TBD

### 5. List Maximum
TBD

## Author
Ferdinandus Richard
