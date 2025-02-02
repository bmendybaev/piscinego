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

	// Print x value (42)
	printNumber(points.x)

	// Print ", y = "
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	// Print y value (21)
	printNumber(points.y)

	// Print newline
	z01.PrintRune('\n')
}

func printNumber(n int) {
	if n > 9 {
		z01.PrintRune('0' + rune(n/10)) // First digit
	}
	z01.PrintRune('0' + rune(n%10)) // Second digit
}
