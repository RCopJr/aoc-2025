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
*/

func main() {
	input := getInputString("i1.txt")
	banks := strings.Split(input, "\n")
	totalJoltage := 0;
	for _, bank := range banks {
		p1 := 0
		p2 := 1
		joltage := 0
		for p2 < len(bank) {
			//Update joltage if possible
			p1Val := string(bank[p1])
			p2Val := string(bank[p2])
			currJoltage, _ := strconv.Atoi(p1Val + p2Val)
			if (currJoltage > joltage) {
				joltage = currJoltage
			}

			//Update p1 if possible
			p1Int, _ := strconv.Atoi(p1Val)
			p2Int, _ := strconv.Atoi(p2Val)
			if (p2Int > p1Int) {
				p1 = p2
			}

			//Increment p2
			p2 += 1
		}
		totalJoltage += joltage
	}
	fmt.Println("Output:", totalJoltage)
}
