package main

import "github.com/01-edu/z01"

// Define the struct point
type point struct {
	x int
	y int
}

// Function to set values in the struct
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// Function to print numbers without fmt
func printNumber(n int) {
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

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	points := &point{}

	setPoint(points)

	printString("x = ")
	printNumber(points.x)
	printString(", y = ")
	printNumber(points.y)
	z01.PrintRune('\n')
}
