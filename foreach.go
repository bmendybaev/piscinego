package piscine

import "github.com/01-edu/z01"

// ForEach applies function f to each element of slice a
func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

// PrintNbr prints a number without using fmt
func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	if n/10 != 0 {
		PrintNbr(n / 10)
	}

	z01.PrintRune(rune(n%10 + '0'))
}
