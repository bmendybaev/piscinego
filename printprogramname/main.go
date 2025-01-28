package main

import (
	"fmt"
	"os"
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

	// Print the program name using z01.PrintRune
	for _, r := range programName {
		fmt.Printf("%c", r)
	}
	fmt.Println()
}