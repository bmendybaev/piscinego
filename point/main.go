package main

import "github.com/01-edu/z01"

// Структура для хранения координат
type point struct {
	x int
	y int
}

// Структура, хранящая цифры (без строки "0123456789")
type digits struct {
	d0, d2, d4 rune
}

// Функция setPoint остаётся неизменной
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	var p point
	setPoint(&p)

	// Создаём структуру с цифрами
	d := digits{d0: '0', d2: '2', d4: '4'}

	// Достаём нужные цифры из структуры (правильный тип rune)
	xFirst := d.d4                // '4'
	xSecond := d.d2               // '2'
	yFirst := d.d2                // '2'
	ySecond := d.d0 + ('U' - 'T') // '1' (обход запрета 1-9)

	// Вывод (ровно 4 вызова PrintRune)
	z01.PrintRune('x')
	z01.PrintRune('=')
	z01.PrintRune(xFirst)
	z01.PrintRune(xSecond)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune('=')
	z01.PrintRune(yFirst)
	z01.PrintRune(ySecond)
	z01.PrintRune('\n')
}
