package piscine

import "github.com/01-edu/z01"

func JumpOver(str string) string {
	if len(str) < 3 {
		z01.PrintRune('\n')
		return "\n"
	}

	s := ""
	c := 3
	for i := 2; i < len(str); i++ {
		if c == 3 {
			s += string(str[i])
			c = 0
		}
		c++
	}
	s += "\n"

	for _, r := range s {
		z01.PrintRune(r)
	}

	return s
}
