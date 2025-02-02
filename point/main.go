package main

import "github.com/01-edu/z01"

// Структура для хранения значений
type Values struct {
	FortyTwo  int
	TwentyOne int
}

// Структура точки
type point struct {
	x, y int
}

// Функция, изменяющая значения через структуру Values
func setPoint(ptr *point, v Values) {
	ptr.x = v.FortyTwo
	ptr.y = v.TwentyOne
}

// Функция для вывода числа
func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	var digits []rune
	for n > 0 {
		digits = append([]rune{rune(n%10) + '0'}, digits...)
		n /= 10
	}
	for _, d := range digits {
		z01.PrintRune(d)
	}
}

func main() {
	// Создаем экземпляр структуры Values с нужными значениями
	values := Values{
		FortyTwo:  42,
		TwentyOne: 21,
	}

	points := &point{}
	setPoint(points, values)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNumber(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNumber(points.y)
	z01.PrintRune('\n')
}
