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

func printChar(c byte) {
	z01.PrintRune(rune(c))
}

func printStr(s string) {
	for i := 0; i < len(s); i++ {
		printChar(s[i])
	}
}

func printInt(n int) {
	if n == 0 {
		printChar('0')
		return
	}

	if n < 0 {
		printChar('-')
		n = -n
	}

	var rev [10]byte
	idx := 0

	for n > 0 {
		rev[idx] = byte(n%10) + '0'
		n /= 10
		idx++
	}

	for i := idx - 1; i >= 0; i-- {
		printChar(rev[i])
	}
}

func main() {
	points := &point{}

	setPoint(points)

	printStr("x = ")
	printInt(points.x)
	printStr(", y = ")
	printInt(points.y)
	printChar('\n')
}
