package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42 // Change this value as needed
	ptr.y = 21 // Change this value as needed
}

func printDigit(d int) {
	switch d {
	case 0:
		z01.PrintRune('0')
	case 1:
		z01.PrintRune('1')
	case 2:
		z01.PrintRune('2')
	case 3:
		z01.PrintRune('3')
	case 4:
		z01.PrintRune('4')
	case 5:
		z01.PrintRune('5')
	case 6:
		z01.PrintRune('6')
	case 7:
		z01.PrintRune('7')
	case 8:
		z01.PrintRune('8')
	case 9:
		z01.PrintRune('9')
	}
}

func printInt(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		printDigit(n / 10)
	}
	printDigit(n % 10)
}

func main() {
	points := &point{}
	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune('=')
	printInt(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune('=')
	printInt(points.y)
	z01.PrintRune('\n')
}
