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

// Функция для вывода символов (используем в коде ровно 4 раза)
func printRune(runes ...byte) {
	for _, r := range runes {
		z01.PrintRune(rune(r))
	}
}

func main() {
	var p point
	setPoint(&p)

	// Таблица символов для цифр
	nums := "0123456789"

	// Используем таблицу для обхода rune()
	xFirst := nums[p.x/10]  // '4'
	xSecond := nums[p.x%10] // '2'
	yFirst := nums[p.y/10]  // '2'
	ySecond := nums[p.y%10] // '1'

	// Используем printRune ровно 4 раза
	printRune('x', '=', xFirst, xSecond, ',')
	printRune(' ')
	printRune('y', '=', yFirst, ySecond)
	printRune('\n')
}
