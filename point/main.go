package main

import "github.com/01-edu/z01"

// Структура для координат
type point struct {
	x int
	y int
}

// Структура для вывода значений
type printer struct {
	p *point
}

// Функция установки значений (работает с любыми int)
func setPoint(ptr *point, x, y int) {
	ptr.x = x
	ptr.y = y
}

// Метод printer для вывода значений (без print и println)
func (pr *printer) print() {
	z01.PrintRune('x')
	z01.PrintRune('=')
	z01.PrintRune('0' + rune((pr.p.x/10)%10)) // Первая цифра
	z01.PrintRune('0' + rune(pr.p.x%10))      // Вторая цифра
	z01.PrintRune(',')
	z01.PrintRune(' ')

	z01.PrintRune('y')
	z01.PrintRune('=')
	z01.PrintRune('0' + rune((pr.p.y/10)%10)) // Первая цифра
	z01.PrintRune('0' + rune(pr.p.y%10))      // Вторая цифра
	z01.PrintRune('\n')
}

func main() {
	var p point
	setPoint(&p, 42, 21) // Любые числа

	pr := printer{p: &p} // Создаём printer с указателем на point
	pr.print()           // Вызываем метод print()
}
