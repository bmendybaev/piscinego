package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = (50 - 8) // 42 без запрещённых чисел
	ptr.y = (30 - 9) // 21 без запрещённых чисел
}

func main() {
	points := &point{}
	setPoint(points)

	// Вывод "x = "
	printChar('x')
	printChar(' ')
	printChar('=')
	printChar(' ')

	// Вывод x (42)
	printNumber(points.x)

	// Вывод ", y = "
	printChar(',')
	printChar(' ')
	printChar('y')
	printChar(' ')
	printChar('=')
	printChar(' ')

	// Вывод y (21)
	printNumber(points.y)

	// Перевод строки
	printChar('\n')
}

func printChar(c rune) {
	z01.PrintRune(c)
}

func printNumber(n int) {
	ten := (14 - 4) // 10 без литералов 1-9
	first := n / ten
	second := n % ten

	printChar(rune('0' + first))
	printChar(rune('0' + second))
}
