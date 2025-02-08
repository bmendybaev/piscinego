package piscine

import "github.com/01-edu/z01"

func JumpOver(str string) string {
	if len(str) < 3 {
		z01.PrintRune('\n')
		return "\n"
	}

	s := ""
	for i := 2; i < len(str); i += 3 {
		s += string(str[i])
	}

	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')

	return s + "\n"
}
