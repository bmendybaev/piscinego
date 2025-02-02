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

func printNumber(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	div := 1
	for temp := n; temp >= 10; temp /= 10 {
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

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNumber(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')

	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNumber(points.y)
	z01.PrintRune('\n')
}
