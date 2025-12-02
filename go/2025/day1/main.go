/*
Input:
- List of ID ranges (i.e. 9-100 would be the list of ids from 9 to 100 inclusive)
- Comma separated
- Seems like input would always go lower-higher
- Numbers will never have leading 0's

Output:
- Want to get the total sum of all invalid ids within all the ranges given
- Invalid Id: a number which is made up of some sequence of digits twice
	- 1010 -> invalid, 11111 -> valid, 456456 -> invalid, 468 -> valid
	- ID must have even digits
	- Must only repeat twice

Naive Approach:
- Loop though each id range from first number to last number
- For each number, you split it in half
- compare the first half with second half, if equal -> add to total sum
- Time: O(n*m*a) -> n is number of ranges, m is size of range, a is size of number in range
- Space: O(m) -> m is size of range (Maybe to perform the comparison)

Improved Approach
- We know only even digit numbers need to be checked so can skip all odd digit numbers
- May not need to actually loop through the number itself to check if it is invalid
- Instead of incrementing, then comparing first and second half, we already know what the invalid version
of an ID would be given its last half digits
- i.e. if I am at ID 1151, I know that the next ID in my range is 1552. Instead of checking if this is invalid, I already know what the 
invalid version of this number would have to be (5252). if that number is in the range of my ids, then it is one of the invalid ids.
- nevermind not really an optimization tbh

Example:
1000: valid
- We know that only one specific combination of last half digits will make this a invalid id -> 10
- All ids inbetween this current id and that one are valid
1010: the closest invalid id to my number -> can skip all of those other ones
- If this id is in range, you add to sum
- Now you have found the invalid id for numbers starting with 10. The only possible next invalid id in this 
would be for numbers starting in 11.
- If that number is no longer in range, there are no other invalid ids. If it is in range, check its invalid counterpart again and continue
- Time: O(n*m) 
- Space: O(1) 

Approach:
- Loop through each range
- Keep track of start and end of range
- Keep track of current number half number
- Keep track of total sum
- Start loop until no more invalid numbers are present
	- Create current number invalid version based on first half of digits
	- If number is in range, add that number to sum
	- If number is not in range, break out of this loop and go into the next range
	- New current number half equals last half number + 1
	- Continue from beginning

Corner Cases:
- Odd digit numbers: 
	- if at 101 for example, next possible invalid id has to be 1010 so it is okay to skip?
	- Bigger example 12541, next possible invalid has to be 125125
	- There example 95-115
		- 9 -> next invalid has to be 99
		- 99 -> what to do in this case? Maybe add 1 to the half number and continue if number is in range?
		- 100 -> next invalid has to be 1010. Number is not in range, so we know that there are no more invalid in this range
	- From looking at their examples this should pretty much work
- What if number is like 99 or something
	- if 9 -> 1
	- if 99 -> 10
	- If half number ends in 0, take off that last 0 and continue -> or just check if length changes
	- Once half number exceeds max half number you stop
- Range starts below 10
	- if end is less than or equal to 10 can just skip
	- just push range to 10
*/
package main

import (
	"os"
	"path/filepath"
	"fmt"
	"strings"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}	
}

func getInputString(filename string) string {
	cwd, err := os.Getwd()
	check(err)
	path := filepath.Join(cwd, filename);
	input, err := os.ReadFile(path)
	check(err)
	return strings.TrimSpace(string(input))
}

func main() {
	input := getInputString("i1.txt");
	intervals := strings.Split(input, ",")

	total := 0

	for _, interval := range intervals {
		boundaries := strings.Split(interval, "-")
		currString := boundaries[0]

		end, err := strconv.Atoi(boundaries[1])
		check(err)

		for {
			curr, _ := strconv.Atoi(currString)

			//Check if half-half
			isInvalid := currString[:len(currString) / 2] == currString[len(currString) / 2:]
			// fmt.Println(end, currString[:len(currString) / 2], currString[len(currString) / 2:])
			
			//If half-half, add to sum
			if (isInvalid) {
				total += curr
				fmt.Println(currString)
			}

			//Increment
			curr += 1
			currString = strconv.Itoa(curr)

			//if start > end, break
			if (curr > end) {
				break;
			}
		}
	}
	fmt.Println(total)
}


// func main() {
// 	input := getInputString("itest.txt");
// 	intervals := strings.Split(input, ",")
//
// 	total := 0
//
// 	for _, interval := range intervals {
// 		boundaries := strings.Split(interval, "-")
//
// 		startString := boundaries[0]
// 		half := (len(startString) + 1) / 2
// 		currHalfString := startString[:half]
//
// 		start, err := strconv.Atoi(boundaries[0])
// 		check(err)
//
// 		endString := boundaries[1]
// 		endHalf := (len(endString) + 1) / 2
// 		maxHalfString := endString[:endHalf]
// 		maxHalf, err := strconv.Atoi(maxHalfString)
// 		check(err)
//
// 		end, err := strconv.Atoi(boundaries[1])
// 		check(err)
//
// 		for {
// 			nextInvalidIdString := currHalfString + currHalfString
// 			nextInvalidId, err := strconv.Atoi(nextInvalidIdString)
// 			check(err)
//
// 			if nextInvalidId >= start && nextInvalidId <= end {
// 				total += nextInvalidId
// 				check(err)
// 				fmt.Println("invalid id:", nextInvalidId)
// 			}
//
// 			lenCurrHalf := len(currHalfString)
// 			currHalf, err := strconv.Atoi(currHalfString)
// 			fmt.Println(currHalf, maxHalf, endString)
// 			currHalf += 1
// 			currHalfString = strconv.Itoa(currHalf)
//
// 			if len(currHalfString) > lenCurrHalf {
// 				currHalf /= 10
// 			}
//
// 			if (currHalf > maxHalf) {
// 				break;
// 			}
//
// 			currHalfString = strconv.Itoa(currHalf)
// 		}
// 	}
// }






