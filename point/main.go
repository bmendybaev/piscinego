package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42 // Change this value as needed
	ptr.y = 21 // Change this value as needed
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

	digits := [10]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	div := 1
	for i := n; i > 9; i /= 10 {
		div *= 10
	}

	for div > 0 {
		digit := n / div
		z01.PrintRune(digits[digit])
		n %= div
		div /= 10
	}
}

func main() {
	points := &point{}
	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune('=')
	printInt(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune('=')
	printInt(points.y)
	z01.PrintRune('\n')
}
