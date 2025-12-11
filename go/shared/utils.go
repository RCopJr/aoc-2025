package utils

import (
	"os"
	"path/filepath"
	"strings"
	"fmt"
)

//NOTE: If func starts with a capital letter it is automatically exported

func Check(e error) {
	if e != nil {
		panic(e)
	}	
}

func GetInputString(filename string) string {
	cwd, err := os.Getwd()
	Check(err)
	path := filepath.Join(cwd, filename);
	input, err := os.ReadFile(path)
	Check(err)
	return strings.TrimSpace(string(input))
}

func BuildByteGrid(input string) [][]byte {
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

func PrintGrid(grid [][]byte) {
	for r := range len(grid) {
		for c := range len(grid[0]) {
			fmt.Print(string(grid[r][c]))
		}
			fmt.Print("\n")
	}
}
