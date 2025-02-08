package piscine

import "github.com/01-edu/z01"

func JumpOver(str string) string {
	if len(str) < 3 {
		z01.PrintRune('\n')
		return "\n"
	}

	result := ""
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	if result == "" {
		z01.PrintRune('\n')
		return "\n"
	}

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')

	return result + "\n"
}
