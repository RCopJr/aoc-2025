package utils

import (
	"os"
	"path/filepath"
	"strings"
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


