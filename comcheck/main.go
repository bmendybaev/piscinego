package main

import (
	"fmt"
	"os"
)

func main() {
	// List of strings to match
	targets := map[string]bool{
		"01":        true,
		"galaxy":    true,
		"galaxy 01": true,
	}

	// Iterate over command-line arguments
	for _, arg := range os.Args[1:] {
		if targets[arg] {
			fmt.Println("Alert!!!")
			return
		}
	}
}
