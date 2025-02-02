package main

import "github.com/01-edu/z01"

// Структура точки
type point struct {
	x, y int
}

// Функция для установки значений (оставляем 42 и 21)
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// Функция для печати числа
func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var digits [10]rune
	for i := range digits {
		digits[i] = '0' + rune(i)
	}
	var result []rune
	for n > 0 {
		result = append([]rune{digits[n%10]}, result...)
		n /= 10
	}
	for _, d := range result {
		z01.PrintRune(d)
	}
}

func main() {
	points := &point{}
	setPoint(points) // Значения останутся 42 и 21

	// Выводим результат
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printInt(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printInt(points.y)
	z01.PrintRune('\n')
}
