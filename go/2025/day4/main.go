package main;

import (
	"fmt"
	"strings"
	"slices"
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

func printGrid(grid [][]byte) {
	for r := range len(grid) {
		for c := range len(grid[0]) {
			fmt.Print(string(grid[r][c]))
		}
			fmt.Print("\n")
	}
}

/*
New Idea:
- once you remove a roll, you can now re consider without that role being there anymore
Naive Approach:
- Just use recursion on the new grid until your valid rolls is 0 then just return your total valid rolls?
Slight Improvement with Tradeoff
- Add some space complexity but really reduce the amount of searches you are doing
- The idea is that after you remove rolls, the only ones you really need to check are any rolls adjacent to the removed ones. These
Will be the only rolls that will have changing adjacent numbers

New Alogorithm:
- Use recursion
- Make sure that search list is all values at the beginning
- Base case: if list of search is empty return 0
- Search through grid and keep track of number of valid rolls.
- Also keep track of coordinates that were removed
- Also keep track of nodes adjacent to those removed coordinates
- Before calling other function make sure to looop through coordinates that were removed and turn them into . so that next
search works correctly
- Return the number of valid rolls + the number of valid rolls that are returned by that new grid and nodes to visit
Corner Case:
- The rolls list will have duplicate values. So it is doing it twice for a coordinate before making it .
- On the next run, a roll may need to be removed, but if the rolls array has it twice, it will remove it twice
- Can either just count difference in @ to . or check if roll is in removed list 
*/

var directions = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
	{1, 1},
}

func searchGrid(grid [][]byte, rolls [][2]int, numRows int, numCols int) int {
	if len(rolls) == 0 {
		printGrid(grid)
		return 0
	}
	validRolls := 0
	removedRolls := [][2]int{}
	newRolls := [][2]int{}
	for _, coord := range rolls {
			r := coord[0]
			c := coord[1]
			if slices.Contains(removedRolls, coord) || grid[r][c] == '.' {
				continue	
			}
			numAdj := 0
			adjRolls := [][2]int{}
			for _, direction := range directions {
				newR := r + direction[0]
				newC := c + direction[1]
				if (newR >= 0 && newR < numRows) && (newC >= 0 && newC < numCols) && grid[newR][newC] == '@' {
					adjRolls = append(adjRolls, [2]int{newR, newC})
					numAdj += 1
				} 
			}
			if numAdj < 4 {
				validRolls += 1 //NOTE: corner case where adding 1 when not supposed to
				removedRolls = append(removedRolls, [2]int{r, c})
				newRolls = append(newRolls, adjRolls...)
			}
	}
	
	for _, roll := range removedRolls {
		grid[roll[0]][roll[1]] = '.'
	}

	return validRolls + searchGrid(grid, newRolls, numRows, numCols)
}

func part2() {
	input := utils.GetInputString("1.txt")	
	grid := buildByteGrid(input)
	numRows := len(grid)
	numCols := len(grid[0])

	//Make initial rolls just every single coordinate in grid
	rolls := [][2]int{}
	for r := range numRows {
		for c := range numCols {
			rolls = append(rolls, [2]int{r, c})
		}
	}
	
	validRolls := searchGrid(grid, rolls, numRows, numCols)
	fmt.Println(validRolls)
}

func part1() {
	input := utils.GetInputString("1.txt")	
	grid := buildByteGrid(input)
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
	// part1()
	part2()
}
