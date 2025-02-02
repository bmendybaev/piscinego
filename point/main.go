package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printPoint(p *point) {
	printString("x = ")
	printInt(p.x)
	printString(", y = ")
	printInt(p.y)
	z01.PrintRune('\n')
}

func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	div := 1
	for i := n; i > 9; i /= 10 {
		div *= 10
	}

	for div > 0 {
		digit := n / div
		z01.PrintRune('0' + rune(digit))
		n %= div
		div /= 10
	}
}

func main() {
	points := &point{}
	setPoint(points)
	printPoint(points)
}
