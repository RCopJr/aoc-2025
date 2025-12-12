package main

import (
	"aoc/shared"
	"fmt"
)

/*
Input:
- Starting point where beam starts going downwards
- A bunch of splitters
- At each splitter, the beam splits into two beams going downwards at c - 1 and c + 1
Output:
- At the of the downward journay, how many beams hit the final row
Naive Approach:
- Can just make an array that stores whether there is a beam in a specific column
- Find the starting point and set that value in the column to 1 or true
- Iterate downwards through the rows
- Check all true columns
- If there is a splitter at a true column, remove that true column and add the two columns beside it
- If you are adding a column and it is already true, its whatever
- Time: O(r*c) -> For every row you are checking every column
- Space: O(c) -> need to store entire column of values for each possible beam
Algorithm:
- Create array of size column with all false values
- Start iterating through the rows
- ...
Improved Approach?
- Maybe where we dont need a space complexity?
- Mabe recursion?
	- Seems like at each splitter you are doing the same operation
	- Seems hard to implement since answer needs the ones that make it to the end
- lets just go with naive approach for now
- Wait, the only thing you need to check is the start and each splitter?
- Rows that have no splitters can be ignored
- Why don't we just edit the input itself?
- Remove the empty rows.
- For every splitter make the character beside it a |
- On the next row for every splitter, if character ontop of it is a |, just do the same
- In the final row, count how many | there are
Misread:
- Problem is how many times does it split
- So yea, just go down each row, draw line beside splitter if line above splitter.
- Everytime you do this, just increase output by 1
- Also seems like every important line is just the odd lines
Corner Thing:
- Solution doesnt work because beams must keep giong passed the first row
- Just keep track of current beams
Other solution:
- So just to remember: Just want number of splits + beams keep going unless they are split
- Just keep track of current beams with an array. if you run into any splitters and on a current beam, update array andn increase output by one
- Can still go row by row
*/

func part1(input string) {
	grid := utils.BuildByteGrid(input)
	numRows := len(grid)
	numCols := len(grid[0])
	beams := make([]bool, numCols)
	output := 0

	for r := range numRows {
		for c := range numCols {
			val := grid[r][c]
			if val == 'S' {
				beams[c] = true
			} else if val == '^' && beams[c] == true {
				output += 1
				if c-1 >= 0 {
					beams[c-1] = true
				}
				if c+1 < numCols {
					beams[c+1] = true
				}
				beams[c] = false
			}
		}
	}

	utils.PrintGrid(grid)
	fmt.Println(output)
}

/*
New Changes:
- Want the number of possible paths the beam can go on starting at S
- Maybe think of it recursively?
- At a splitter the amount of timelines = the amount of timelines of left path + timelines of right path
- Base case: If you reach the end, you know that you only have one timeline
- If you reach a splitter, your number of time lines = left timelines + right timelines
- If you reach a ., you just continue going down
- Time: should def be smaller than just trying all of the paths until you reach a leaf node
	- Once you reach a leaf node, you just go backwards since you know the answers now
Recursive algorithm:
- Call function where S starts. Just needs the r and c value
- While grid value == '.'  and in bounds just increment r
- If value at coord == out of bounds, return 1
- Else (means splitter) return function at (r, c - 1)  + function at (r, c + 1)
Can we use dynamic programming?
- Since multiple beams can hit the same splitter, might be better to store the splitter value before it returns somewhere, That way future beams dont need to traverse the rest of the tree again
- Time: O(n). With DP would only traverse past a splitter once
- Space: O(n). Since DP would worst case need an array for every possible coordinate
Input Notes:
- No splitters on the edge, so no need to check column out of bounds
- No splitters beside each other so can assume '.' right after function call
*/

type Coord struct {
	R int
	C int
}

func part2(input string) {
	grid := utils.BuildByteGrid(input)
	timelines := map[Coord]int{}
	numRows := len(grid)

	var getTimelines func(r int, c int) int
	getTimelines = func(r int, c int) int {
		coord := Coord{r, c}
		val := grid[r][c]
		for r < numRows && val == '.' {
			r += 1
			if r == numRows {
				break
			}
			val = grid[r][c]
		}

		if r >= numRows {
			return 1
		} else {
			var numTimelines int
			if _, exists := timelines[coord]; exists {

				return timelines[coord]
			}
			numTimelines = getTimelines(r, c - 1) + getTimelines(r, c + 1)
			timelines[coord] = numTimelines
			return numTimelines
		}
	}

	for c := range len(grid[0]) {
		if grid[0][c] == 'S' {
			output := getTimelines(1, c) //NOTE: Start at r = 1 to avoid 'S' edge case
			fmt.Println(output)
		}
	}

}

func main() {
	input := utils.GetInputString("actual.txt")
	part2(input)
}
