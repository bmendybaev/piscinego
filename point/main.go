package main

import "github.com/01-edu/z01"

// Структура точки
type point struct {
	x, y int
}

// Функция установки значений (работает с int)
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points) // Можно менять points.x и points.y

	// Вывод результата
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	// Печать десятков
	if points.x/10 == 4 {
		z01.PrintRune('4')
	} else if points.x/10 == 2 {
		z01.PrintRune('2')
	}

	// Печать единиц
	if points.x%10 == 2 {
		z01.PrintRune('2')
	} else if points.x%10 == 1 {
		z01.PrintRune('1')
	}

	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	// Печать десятков
	if points.y/10 == 4 {
		z01.PrintRune('4')
	} else if points.y/10 == 2 {
		z01.PrintRune('2')
	}

	// Печать единиц
	if points.y%10 == 2 {
		z01.PrintRune('2')
	} else if points.y%10 == 1 {
		z01.PrintRune('1')
	}

	z01.PrintRune('\n')
}
