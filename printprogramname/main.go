package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get the name of the program from os.Args[0]
	programName := os.Args[0]

	// Extract only the base name without the full path
	for i := len(programName) - 1; i >= 0; i-- {
		if programName[i] == '/' || programName[i] == '\\' {
			programName = programName[i+1:]
			break
		}
	}

	// Print each character using z01.PrintRune
	for _, r := range programName {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n') // Print a newline at the end
}
