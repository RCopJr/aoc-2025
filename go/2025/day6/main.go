package main

import (
	"aoc/shared"
	"fmt"
	"strconv"
	"strings"
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

	for groupIdx, part := range parts[:len(parts)-1] {
		rowNums := strings.Fields(part)
		for idx, num := range rowNums {
			numInt, _ := strconv.Atoi(num)
			nums[idx][groupIdx] = numInt
		}
	}

	operators := strings.Fields(parts[len(parts)-1])
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

/*
Changes:
- The numbers are read differently
- You no read numbers from top to bottom right to left
- Spaces cant be fully disregarded now since they play a role in where the numbers are positioned
- Important: Each group separated by a full column of just spaces (This is how you know to move on)
Approach:
- Lets just create a 2D array with spaces included
- Keep track of the final row of operators to know what the operators are
- Then start by looping downwards across the rows to create lists of numbers for that group
- Once a column of just spaces is hit, you apply the arithmetic to your list of numbers based on the current operator
- Then you update the operator and start over
Algorithm:
- Create list with spaces included for first 4 rows
- Create list of operators from last row
- Keep track of current operator index
- Loop through each index in number of rows
	- Loop through the columns at this index
	- For each number in column, create string concatenating them together (if space, just skip)
	- Once you reach the bottom, update the index
	- If you reach the bottom and there are no numbers, apply arithmetic to current string of numbers, and update operator index
		- Make sure to reset the group after the arithmetic is finished
- Still a linear time algorithm
Input Notes:
- All rows are the same length
*/

func part2(input string) {
	parts := strings.Split(input, "\n")
	operators := strings.Fields(parts[len(parts)-1])
	operatorIdx := 0
	nums := parts[:len(parts)-1] //NOTE: Want to ignore final row
	cols := len(parts[0])
	rows := len(parts) - 1
	output := 0

	group := []int{}
	for c := range cols {
		digits := []byte{}
		for r := range rows {
			digit := nums[r][c]
			if digit == ' ' {
				continue
			} else {
				digits = append(digits, digit)
			}
		}
		if len(digits) != 0 {
			group = append(group, bytesToInt(digits))
		} else {
			output += evaluateGroup(group, operators[operatorIdx])
			operatorIdx += 1
			group = []int{}
		}
	}

	//NOTE: Corner case: if end is reached and still have number in group should also add to output
	output += evaluateGroup(group, operators[operatorIdx])

	fmt.Println(output)
}

func evaluateGroup(group []int, operator string) int {
	acc := 0
	if operator == "+" {
		for _, num := range group {
			acc += num
		}
	} else {
		acc = 1
		for _, num := range group {
			acc *= num
		}
	}
	return acc
}

// NOTE: Could use strconv.Atoi() but this seemed cool
func bytesToInt(b []byte) int {
	n := 0
	for _, c := range b {
		n = n*10 + int(c-'0')
	}
	return n
}

func main() {
	input := utils.GetInputString("1.txt")
	part2(input)
}
