package main  


import (
	"os"
	"path/filepath"
	"strings"
	"strconv"
	"fmt"
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

/*
Input:
- banks of batteries
- each row is a bank, each element in that row is a battery
Output:
- for each row, want to produce the biggest number
- Numbers are produced by taking 2 batteries in their respective order and create a 2 digit number with them
- You want to find the biggest possible number for each bank
- Actual output is total sum of each bank
Naive Approach:
- For each number, check all the numbers to the right of it and generate a number
- Keep track of the max number
- Time: O(n^2)
Improved Approach:
- Two pointer approach
- Since you know that first pointer should always be before second. Keep iterating second pointer until you find a number
bigger than the first pointer. That becomes your new first pointer
bigger than the first pointer. That becomes your new first pointer
- As you do this iteration, you keep track of the numbers you are creating and keep the max
- This works because the only time a number can for sure be greater than a number starting at p1 will be if p2 is bigger
- Time: O(n)
Possible Vorner Cases:
- if there is only one element maybe?

Part 2:
- instead of 2 digits you have 12 digits
- You want to find the group of 12 positions that go from most biggest digit from left to right if possible
- You know that your first digit must be before the final 12 digits -> Maybe we can start there?
New Algorithm
- Start at the end with all digits
- Make second pointer keep on traversing backwards
- Have 11 pointers for other spots
- Find max number between 1 and 2 and put 2 there, put 2 at most left biggest number between 1 and 2
- Do the same thing with 2 and 3 and so on...
- Do the same thing with 2 and 3 and so on...
Other Algorithm:
- If you start at the end, cant you just instantly send the first digit to the biggest digit to its left?
- Your gonna want to send it to the most left of the biggest so your second digit has the most possible options
Steps:
- Setup all 12 pointers at the end of the string
- Move the first pointer the biggest number to its left (only way this number can get bigger)
- Move second number to the biggest number to its left between itself and 1 (Only way this number can now get bigger)
- Do the same for all 12 digits
- Can pointer be an array of indices?
*/

func main() {
	input := getInputString("i1.txt")
	banks := strings.Split(input, "\n")
	totalJoltage := 0;
	// pointers := [12]int{88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}
	// pointers := [12]int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	for _, bank := range banks {
		pointers := [12]int{88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}
		start := 0
		for idx, pointer := range pointers {
			//Find index of biggest value to the left of pointer (inclusive) and set pointer to that value
			maxIdx := pointer
			maxVal, _ := strconv.Atoi(string(bank[maxIdx]))
			fmt.Println("pointer and start before for loop:", pointer, start)
			for newPointer := pointer; newPointer >= start; newPointer-- {
				// fmt.Println("pointers within for loop", pointer, newPointer)
				pointerVal, _ := strconv.Atoi(string(bank[newPointer]))
				if (pointerVal >= maxVal) { //Want >= since we want it to go to the most left option
					maxIdx = newPointer
					maxVal = pointerVal
				}
			}
			pointers[idx] = maxIdx
			//Update start to current pointer at end of this loop (Also update the start val)
			start = maxIdx + 1 //Want to make start the index after since you dont want to duplicate indices
		}
		fmt.Println(pointers)
		joltageString := ""
		for _, pointer := range pointers {
			joltageString += string(bank[pointer])
		}
		fmt.Println(joltageString)
		joltage, _ := strconv.Atoi(joltageString)
		fmt.Println(joltage)
		totalJoltage += joltage
		fmt.Print("\n")
	}
	fmt.Println("Output:", totalJoltage)
}

// func main() {
// 	input := getInputString("i1.txt")
// 	banks := strings.Split(input, "\n")
// 	totalJoltage := 0;
// 	for _, bank := range banks {
// 		p1 := 0
// 		p2 := 1
// 		joltage := 0
// 		for p2 < len(bank) {
// 			//Update joltage if possible
// 			p1Val := string(bank[p1])
// 			p2Val := string(bank[p2])
// 			currJoltage, _ := strconv.Atoi(p1Val + p2Val)
// 			if (currJoltage > joltage) {
// 				joltage = currJoltage
// 			}
//
// 			//Update p1 if possible
// 			p1Int, _ := strconv.Atoi(p1Val)
// 			p2Int, _ := strconv.Atoi(p2Val)
// 			if (p2Int > p1Int) {
// 				p1 = p2
// 			}
//
// 			//Increment p2
// 			p2 += 1
// 		}
// 		totalJoltage += joltage
// 	}
// 	fmt.Println("Output:", totalJoltage)
// }
