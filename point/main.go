package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = '(' * 6 // 42
	ptr.y = '!' * 3 // 21
}

func printNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var digits []rune
	sign := false
	if n < 0 {
		sign = true
		n = -n
	}
	for n > 0 {
		digits = append(digits, rune(n%('0'+'1'))+'0')
		n /= ('0' + '1')
	}
	if sign {
		z01.PrintRune('-')
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	points := &point{}

	setPoint(points)

	printStr("x = ")
	printNbr(points.x)
	printStr(", y = ")
	printNbr(points.y)
	z01.PrintRune('\n')
}
