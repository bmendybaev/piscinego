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

	// Workaround: Convert numbers manually using a slice
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	// Print x value (42)
	z01.PrintRune([]rune(digits[points.x/10])[0])
	z01.PrintRune([]rune(digits[points.x%10])[0])

	// Print ", y = "
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	// Print y value (21)
	z01.PrintRune([]rune(digits[points.y/10])[0])
	z01.PrintRune([]rune(digits[points.y%10])[0])

	// Print newline
	z01.PrintRune('\n')
}
