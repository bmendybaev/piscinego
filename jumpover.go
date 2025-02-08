package piscine

import (
	"github.com/01-edu/z01"
)

func JumpOver(str string) string {
	result := ""

	// Iterate over the string and pick every third character
	for i, char := range str {
		if (i+1)%3 == 0 { // (i+1) to make it 1-based indexing
			result += string(char)
		}
	}

	// If result is empty, return newline
	if result == "" {
		return "\n"
	}

	// Append newline at the end
	return result + "\n"
}
