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

	// Print first number (42)
	if points.x/10 > 0 {
		z01.PrintRune('0' + rune(points.x/10))
	}
	z01.PrintRune('0' + rune(points.x%10))

	// Print ", y = "
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	// Print second number (21)
	if points.y/10 > 0 {
		z01.PrintRune('0' + rune(points.y/10))
	}
	z01.PrintRune('0' + rune(points.y%10))

	// Print newline
	z01.PrintRune('\n')
}
