package main

import (
	"aoc/shared"
	"fmt"
	"strconv"
	"strings"
	"slices"
	"cmp"
)

/*
Input:
- list of ranges followed by list of ids
Output:
- Number of ids that are within the ranges
Naive Approach:
- just go through each range for each id and check if you go through all the ranges without finding one.
- if you find just add one to answer and break
Improved Approach:
- Can maybe do intervals not hit once then you would exit the loops earlier, but time complexity stays the same
*/

func part1(input string) {
	parts := strings.Split(input, "\n\n")
	ranges := parts[0]
	ids := parts[1]
	rangeList := strings.Split(ranges, "\n")
	idList := strings.Split(ids, "\n")
	validIntervals := [][2]int{}

	for _, interval := range rangeList {
		endpoints := strings.Split(interval, "-")
		start, _ := strconv.Atoi(endpoints[0])
		end, _ := strconv.Atoi(endpoints[1])
		validIntervals = append(validIntervals, [2]int{start, end})
	}

	output := 0
	for _, id := range idList {
		idInt, _ := strconv.Atoi(id)
		isValid := false
		for _, interval := range validIntervals {
			if idInt >= interval[0] && idInt <= interval[1] {
				isValid = true
				break
			}
		}
		if isValid {
			output += 1
		}
	}

	fmt.Println(output)
}
/*
Update:
- Now we just want all possible valid ids based on the ranges given.
Naive Approach:
- go through every range, append number to valid ids, if number is already in valid ids, just skip it
- Time: O(n^2*m^2) -> something like that since for every range, you check al numbers, and search all already checked numbers
Improved Approach:
- Treat this as an interval problem
- sort all of the ranges by start
- Loop through all of the ranges
	- For first interval just append to output
	- If second interval start is greater than first interval end, append that to the list
	- if second interval start is lesser than first interval end, check to see if first interval end greater or lesser than, 
	second interval end
	- Make the end of the most recent interval equal to the bigger of the two
	- Continue until no more in list
- Return the total length of all of the intervals in the final output array
- Note: To make it even better, dont make a new array with all of the intervals. 
	- just keep track of current interval.
	- if newest interval is within current, just update end. 
	- if newest interval is outside current, add to output the end minus the start of current interval.
	Then make the newest interval equal to the current
	- Time: O(n*logn) + O(n)
*/
func part2(input string) {
	parts := strings.Split(input, "\n\n")
	ranges := parts[0]
	intervalStrings := strings.Split(ranges, "\n")
	intervals := [][2]int{}

	for _, intervalString := range intervalStrings {
		endpoints := strings.Split(intervalString, "-")
		start, _ := strconv.Atoi(endpoints[0])
		end, _ := strconv.Atoi(endpoints[1])
		intervals = append(intervals, [2]int{start, end})
	}

	startCmp := func (a [2]int, b[2]int) int {
		return cmp.Compare(a[0], b[0])
	}

	slices.SortFunc(intervals, startCmp)

	output := 0
	currInterval := [2]int{}

	for i, interval := range intervals {
		if i == 0 {
			currInterval = interval
			continue
		}

		if interval[0] <= currInterval[1] {
			currInterval[1] = max(interval[1], currInterval[1])
		} else {
			output += currInterval[1] - currInterval[0] + 1
			currInterval = interval
		}
	}

	//NOTE: Corner case: if you have a currInterval still you also need to add to output
	output += currInterval[1] - currInterval[0] + 1

	fmt.Println(output)
}

func main() {
	input := utils.GetInputString("1.txt")
	// part1(input)
	part2(input)
}
