package main

import "github.com/01-edu/z01"

// Структура точки
type point struct {
	x, y int
}

// Функция для установки значений (можно менять x и y)
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// Функция для печати числа (работает с ЛЮБЫМ int)
func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var digits [2]rune
	digits[0] = '0' + rune(n/10) // первая цифра
	digits[1] = '0' + rune(n%10) // вторая цифра

	if n >= 10 {
		z01.PrintRune(digits[0])
	}
	z01.PrintRune(digits[1])
}

// Функция печати точки (x и y теперь можно менять!)
func printPoint(x, y int) {
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNumber(x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNumber(y)
	z01.PrintRune('\n')
}

func main() {
	points := &point{}
	setPoint(points) // Можно менять points.x и points.y

	// Вызываем printPoint(), передавая x и y
	printPoint(points.x, points.y)
}
