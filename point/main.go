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

func printCharSequence(runes ...rune) {
	for _, r := range runes {
		z01.PrintRune(r)
	}
}

func main() {
	points := &point{}
	setPoint(points)

	printCharSequence('x', '=', rune('0'+points.x/10), rune('0'+points.x%10), ',', ' ', 'y', '=', rune('0'+points.y/10), rune('0'+points.y%10), '\n')
}
