/*
Description

    We want you to calculate the sum of squares of given integers, excluding any negatives.
    The first line of the input will be an integer N (1 <= N <= 100), indicating the number of test cases to follow.
    Each of the test cases will consist of a line with an integer X (0 < X <= 100), followed by another line consisting of X number of space-separated integers Yn (-100 <= Yn <= 100).
    For each test case, calculate the sum of squares of the integers, excluding any negatives, and print the calculated sum in the output.
    Note: There should be no output until all the input has been received.
    Note 2: Do not put blank lines between test cases solutions.
    Note 3: Take input from standard input, and output to standard output.

Rules

    Write your solution using Go Programming Language
    Your source code must be a single file (package main)
    Do not use any for statement
    You may only use standard library packages

    */
package main

import (
	"fmt"
)

var results []int
var nbr_of_values int
var nbr_of_tests int

func main() {

	fmt.Scan(&nbr_of_tests)
	tests(&nbr_of_tests)
	index := 0
	printer(&index, results)
}

func tests(nbt *int) {

	fmt.Scan(&nbr_of_values) //get the number of intergers in the next line

	results = append(results, operation(&nbr_of_values))

	*nbt = (*nbt) - 1
	if (*nbt) != 0 { // check if there is still tests to perform
		tests(&nbr_of_tests)

	}
}

func operation(nbv *int) int {
	if (*nbv) == 0 { // we scaned all the intergers of the current line
		return 0

	} else {
		var k int
		fmt.Scan(&k) // scan the next integer

		if k >= 0 { //check for negative numbers
			k = k * k
			*nbv = (*nbv) - 1
			return k + operation(&nbr_of_values)
		} else {
			k = 0
			*nbv = (*nbv) - 1
			return k + operation(&nbr_of_values)
		}
	}
}

// print the values in the same format as in the demo
func printer(i *int, b []int) {

	if (*i) < len(b) { // if the the index is inferiour to length of slice keep recusing

		fmt.Println(b[(*i)])
		*i = (*i) + 1
		printer(i, b)
	}
}
