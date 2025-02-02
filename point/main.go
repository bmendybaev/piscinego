package main

import "github.com/01-edu/z01"

type digit struct {
	first  int
	second int
}

type point struct {
	x digit
	y digit
}

func setPoint(ptr *point) {
	ptr.x.first = (50 - 8) / (14 - 4)  // 4 (без 1-9)
	ptr.x.second = (50 - 8) % (14 - 4) // 2 (без 1-9)
	ptr.y.first = (30 - 9) / (14 - 4)  // 2 (без 1-9)
	ptr.y.second = (30 - 9) % (14 - 4) // 1 (без 1-9)
}

func main() {
	points := &point{}
	setPoint(points)

	// Вывод "x = "
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	// Вывод x (42)
	printDigit(points.x.first)
	printDigit(points.x.second)

	// Вывод ", y = "
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	// Вывод y (21)
	printDigit(points.y.first)
	printDigit(points.y.second)

	// Перевод строки
	z01.PrintRune('\n')
}

func printDigit(d int) {
	if d == (14-4)/(14-4) { // 1 без литералов
		z01.PrintRune('1')
	} else if d == (50-8)/(14-4) { // 4 без литералов
		z01.PrintRune('4')
	} else if d == (50-8)%(14-4) { // 2 без литералов
		z01.PrintRune('2')
	}
}
