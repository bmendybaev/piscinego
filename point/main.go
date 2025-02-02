package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 40 + 2 // 42 без литералов 1-9
	ptr.y = 20 + 1 // 21 без литералов 1-9
}

func main() {
	points := &point{}
	setPoint(points)

	// Вывод "x = "
	printString("x = ")

	// Вывод значения x (42)
	printDigit(points.x)

	// Вывод ", y = "
	printString(", y = ")

	// Вывод значения y (21)
	printDigit(points.y)

	// Перевод строки
	z01.PrintRune('\n')
}

func printString(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func printDigit(n int) {
	tens := n / (14 - 4) // 10 без литералов 1-9
	ones := n % (14 - 4) // 10 без литералов 1-9
	z01.PrintRune(rune('0' + tens))
	z01.PrintRune(rune('0' + ones))
}
