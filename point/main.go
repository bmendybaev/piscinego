package main

import "github.com/01-edu/z01"

// Define struct point
type point struct {
	x int
	y int
}

// Modify function to accept an int pointer
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// Custom print function using z01.PrintRune
func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	var digits []rune
	for n > 0 {
		digits = append([]rune{rune(n%10) + '0'}, digits...)
		n /= 10
	}

	for _, d := range digits {
		z01.PrintRune(d)
	}
}

func main() {
	points := &point{}
	setPoint(points)

	// Print formatted output without fmt
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printInt(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')

	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printInt(points.y)
	z01.PrintRune('\n')
}
