package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	for first := 99; first >= 0; first-- {
		for second := first - 1; second >= 0; second-- {
			printNumber(first)
			z01.PrintRune(' ')
			printNumber(second)

			if !(first == 1 && second == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}

func printNumber(n int) {
	z01.PrintRune(rune('0' + n/10))
	z01.PrintRune(rune('0' + n%10))
}
