package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printNumber(n int) {
	digits := "0123456789"
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
		z01.PrintRune(rune(digits[n/div]))
		n %= div
		div /= 10
	}
}

func main() {
	points := &point{}

	setPoint(points)

	fmt.Print("x = ")
	printNumber(points.x)
	fmt.Print(", y = ")
	printNumber(points.y)
	fmt.Print("\n")
}
