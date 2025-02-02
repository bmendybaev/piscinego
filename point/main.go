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
	str := "x = "
	for i := 0; i < len(str); i++ {
		z01.PrintRune(rune(str[i]))
	}

	// Convert x (42) to string and print
	xStr := []byte{byte('0' + points.x/10), byte('0' + points.x%10)}
	for i := 0; i < len(xStr); i++ {
		z01.PrintRune(rune(xStr[i]))
	}

	// Print ", y = "
	str = ", y = "
	for i := 0; i < len(str); i++ {
		z01.PrintRune(rune(str[i]))
	}

	// Convert y (21) to string and print
	yStr := []byte{byte('0' + points.y/10), byte('0' + points.y%10)}
	for i := 0; i < len(yStr); i++ {
		z01.PrintRune(rune(yStr[i]))
	}

	// Print newline
	z01.PrintRune('\n')
}
