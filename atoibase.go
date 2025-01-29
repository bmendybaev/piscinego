package piscine

import "github.com/01-edu/z01"

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	baseLen := len(base)
	result := 0

	for i := 0; i < len(s); i++ {
		value := indexInBase(s[i], base)
		if value == -1 {
			return 0
		}
		result = result*baseLen + value
	}

	return result
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}

	return true
}

func indexInBase(c byte, base string) int {
	for i := 0; i < len(base); i++ {
		if base[i] == c {
			return i
		}
	}
	return -1
}

func PrintInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var digits []rune
	for n > 0 {
		digits = append([]rune{rune(n%10) + '0'}, digits...)
		n /= 10
	}

	for _, digit := range digits {
		z01.PrintRune(digit)
	}
	z01.PrintRune('\n')
}
