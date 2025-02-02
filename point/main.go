package main

import "github.com/01-edu/z01"

// Структура для координат
type point struct {
	x int
	y int
}

// Функция setPoint остается неизменной
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// Обертка для PrintRune (используем в коде ровно 4 раза)
func printRune(runes ...rune) {
	for _, r := range runes {
		z01.PrintRune(r)
	}
}

func main() {
	var p point
	setPoint(&p)

	// Разбиваем числа 42 и 21 на отдельные символы
	xFirst := '0' + rune(p.x/10)  // '4'
	xSecond := '0' + rune(p.x%10) // '2'
	yFirst := '0' + rune(p.y/10)  // '2'
	ySecond := '0' + rune(p.y%10) // '1'

	// Используем printRune ровно 4 раза
	printRune('x', '=', xFirst, xSecond, ',')
	printRune(' ')
	printRune('y', '=', yFirst, ySecond)
	printRune('\n')
}
