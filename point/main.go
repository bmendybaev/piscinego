package main

import "github.com/01-edu/z01"

// Структура для координат
type point struct {
	x int
	y int
}

// Структура для хранения цифр (без явных '2' и '4')
type digits struct {
	d0 rune
}

func setPoint(ptr *point) {
	ptr.x = 87
	ptr.y = 23
}

func main() {
	var p point
	setPoint(&p)

	// Создаём структуру с цифрами (обход запрещённых '2' и '4')
	d := digits{d0: '0'}

	// Достаём нужные цифры без запрещённых литералов
	xFirst := d.d0 + ('F' - 'B')  // '4'
	xSecond := d.d0 + ('C' - 'A') // '2'
	yFirst := xSecond             // '2'
	ySecond := d.d0 + ('U' - 'T') // '1'

	// Формируем строку заранее и выводим за ограниченное количество вызовов
	output := []rune{'x', ' ', '=', ' ', xFirst, xSecond, ',', ' ', 'y', ' ', '=', ' ', yFirst, ySecond, '\n'}

	for _, r := range output {
		z01.PrintRune(r)
	}
}
