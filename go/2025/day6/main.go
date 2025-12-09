package main

import (
	"aoc/shared"
	"fmt"
	"strings"
	"strconv"
)

/*
Input:
- List of numbers and operations
- Each line has the same number of x amount of space separated numbers
- On the final line there is a list of space separated operations that are applied to all numbers above it
- The amount of spaces between each number can be arbitrary
Output:
- The total sum of all operations
First approach:
- Get length of line and create 2D array of this size
- Get number of lines to know when at the operator line
- Loop through each line
	- split by space
	- for each value in array, trim, convert to integer, and append to element in this array
	- At the final line break out of this loop
- Loop through each operator
	- Loop through each element in array and apply operator and add to total sum
	- Done
Input Notes:
- There are 5 rows including the operations
- Can use strings.Fields to get just values without spaces
- Only operators are + and *
*/

func part1(input string) {
	parts := strings.Split(input, "\n")

	groupSize := len(parts) - 1
	rowSize := len(strings.Fields(parts[0]))

	nums := make([][]int, rowSize)
	for i := range nums {
		nums[i] = make([]int, groupSize)
	}

	for groupIdx, part := range parts[:len(parts) - 1] {
		rowNums := strings.Fields(part)
		for idx, num := range rowNums {
			numInt, _ := strconv.Atoi(num)
			nums[idx][groupIdx]	= numInt
		}
	}

	operators := strings.Fields(parts[len(parts) - 1])
	output := 0
	for i, problem := range nums {
		operator := operators[i]

		acc := 0
		if operator == "+" {
			for _, num := range problem {
				acc += num
			}
		} else {
			acc = 1
			for _, num := range problem {
				acc *= num
			}
		}
		output += acc
	}

	fmt.Println(output)
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func main() {
	input := utils.GetInputString("1.txt")
	part1(input)
}









