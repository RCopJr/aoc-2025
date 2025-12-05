package main;

import (
	"fmt"
	"strings"
	"aoc/shared"
)

/*
Input:
- grid of room
- @ sign represents a roll of paper
Output:
- The amount of rolls of paper that can be accessed
- a roll of paper can be accessed if it has less than 4 adjacent rolls of paper
- Adjacent means having a roll of paper to the left, top, right, bottom, or any of the 4 diagonals
Naive Approach:
- Loop through every coordinate
- Check all 8 directions (Do arithmetic to coordinate)
- If @ sign in 4 or more of the 8 directions, you increase counter by 1
- Time: O(n*m) -> only 8 directions, so operation per coordinate is constant time
Maybe better approach?
- Could do a graph traversal for each node with one visited list
- As you are keeping track of visited, you check how many times a node has been checked and has been visited
- If this number is 3 or greater, than you know it is not a valid roll of paper
- This approach does seem more complicated and right now dont really need a graph to answer this question
- Would also have a non constant space complexity now

Algorithm:
- Create a 2D matrix for input
- Keep track of total valid rolls
- Create list of direction arithmetic for easier calcualtion
- Create double for loop to go to each coordinate
- At each coordinate, check all 8 directions and keep track of how many have rolls
- if more than 4, stop iteration and increment total sum counter
*/

func buildByteGrid(input string) [][]byte {
	//Define 2D slice
	var grid [][]byte
	//Split input string into rows
	rows := strings.Split(input, "\n")
	//Loop through each line
	for _, row := range rows {
		grid = append(grid, []byte(row))
	}
	return grid
}

func part1() {
	input := utils.GetInputString("1.txt")	
	grid := buildByteGrid(input)
	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
		{1, 1},
	}
	validRolls := 0
	numRows := len(grid)
	numCols := len(grid[0])

	for r := range numRows {
		for c := range numCols {
			if grid[r][c] == '.' {
				continue
			}
			numAdj := 0
			for _, direction := range directions {
				newR := r + direction[0]
				newC := c + direction[1]
				if (newR >= 0 && newR < numRows) && (newC >= 0 && newC < numCols) && grid[newR][newC] == '@' {
					numAdj += 1
				} 
			}
			if numAdj < 4 {
				validRolls += 1
			}
		}
	}
	fmt.Println(validRolls)
}

func main() {
	part1()
}
