package main

import "github.com/01-edu/z01"

// Структура для хранения координат
type point struct {
	x int
	y int
}

// Структура для вывода значений
type printer struct {
	p *point
}

// Устанавливаем любые значения x и y
func setPoint(ptr *point, x, y int) {
	ptr.x = x
	ptr.y = y
}

// Метод printer для вывода значений
func (pr *printer) print() {
	z01.PrintRune('x')
	z01.PrintRune('=')
	z01.PrintRune('0' + rune(pr.p.x/10)) // Первая цифра числа
	z01.PrintRune('0' + rune(pr.p.x%10)) // Вторая цифра числа
	z01.PrintRune(',')
	z01.PrintRune(' ')

	z01.PrintRune('y')
	z01.PrintRune('=')
	z01.PrintRune('0' + rune(pr.p.y/10)) // Первая цифра числа
	z01.PrintRune('0' + rune(pr.p.y%10)) // Вторая цифра числа
	z01.PrintRune('\n')
}

func main() {
	var p point
	setPoint(&p, 42, 21) // Передаём любые числа

	pr := printer{p: &p} // Создаём printer с указателем на point
	pr.print()           // Вызываем метод print()
}
