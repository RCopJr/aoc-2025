package main

import (
	"fmt"
	"aoc/shared"
	"math"
	"strings"
	"strconv"
)

/*
Input:
- list of coordinates of red tiles

Output:
- The area of the biggest rectangle you make by picking two tiles at opposite corners of a rectangle
- Rectangles on the same axis still have an axis of 1
- Seems like you can just do  (abs(x1 - x2) + 1) * (abs(y1 - y2) + 1)

Naive Approach
- For every tile, just get the area of it with every other tile. 
- Find the biggest area, and return those two tiles
- Make sure no duplicates by looping from i in the outer and i + 1 in the inner
- Time: O(n^2)
- Space: O(n)

Improved Approach?
- Was thinking of tackling it in similar way as container with most water, but dont know if it would work
*/

type Coord struct {
	X int
	Y int
}

func getArea(a Coord, b Coord) float64 {
	return (math.Abs(float64(a.X - b.X)) + 1) * (math.Abs(float64(a.Y - b.Y)) + 1)
}

func part1(input string) {
	redTiles := []Coord{}
	for coord := range strings.SplitSeq(input, "\n") {
		tileString := strings.Split(coord, ",")
		x, _ := strconv.Atoi(tileString[0])
		y, _ := strconv.Atoi(tileString[1])
		redTiles = append(redTiles, Coord{X: x, Y: y})
	}

	maxArea := 0.0
	for i := 0; i < len(redTiles) - 1; i++ {
		for j := i + 1; j < len(redTiles); j++ {
			area := getArea(redTiles[i], redTiles[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}

	fmt.Println(int(maxArea))
}

/*
Changes
- There are now green tiles
- Green tiles exist in between any adjacent red tiles in the list
- Adjacent red tiles are either in the same column or row
*/

func part2(input string) {

}

func main() {
	input := utils.GetInputString("actual.txt")
	part1(input)
}
