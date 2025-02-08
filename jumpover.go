package piscine

import "github.com/01-edu/z01"

func JumpOver(str string) string {
	result := ""
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	if result == "" {
		result = "\n"
	} else {
		result += "\n"
	}

	// Print the result using z01.PrintRune
	for _, r := range result {
		z01.PrintRune(r)
	}

	return result
}
