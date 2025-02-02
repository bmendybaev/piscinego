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

func main() {
	points := &point{}
	setPoint(points)

	// Print "x = "
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	// Print x value (42) dynamically
	num := points.x
	div := 1
	for temp := num; temp >= 10; temp /= 10 {
		div *= 10
	}
	for div > 0 {
		digit := num / div
		z01.PrintRune(rune('0' + digit)) // Explicit rune conversion
		num %= div
		div /= 10
	}

	// Print ", y = "
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	// Print y value (21) dynamically
	num = points.y
	div = 1
	for temp := num; temp >= 10; temp /= 10 {
		div *= 10
	}
	for div > 0 {
		digit := num / div
		z01.PrintRune(rune('0' + digit)) // Explicit rune conversion
		num %= div
		div /= 10
	}

	// Print newline
	z01.PrintRune('\n')
}
