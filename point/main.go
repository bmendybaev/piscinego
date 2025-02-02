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

func main() {
	points := &point{}
	setPoint(points)

	printCharSequence('x', '=')
	printInt(points.x)
	printCharSequence(',', ' ', 'y', '=')
	printInt(points.y)
	printCharSequence('\n')
}

func printCharSequence(runes ...rune) {
	for _, r := range runes {
		z01.PrintRune(r)
	}
}

func printInt(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		printCharSequence('0'+rune(n/10), '0'+rune(n%10))
	} else {
		printCharSequence('0' + rune(n))
	}
}
