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

func printRuneSequence(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	points := &point{}
	setPoint(points)

	printRuneSequence("x = ")
	printRuneSequence(toString(points.x))
	printRuneSequence(", y = ")
	printRuneSequence(toString(points.y))
	printRuneSequence("\n")
}

func toString(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""
	if n < 0 {
		result += "-"
		n = -n
	}

	div := 1
	for i := n; i > 9; i /= 10 {
		div *= 10
	}

	for div > 0 {
		digit := n / div
		result += string('0' + digit)
		n %= div
		div /= 10
	}

	return result
}
