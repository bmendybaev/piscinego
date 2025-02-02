package main

import "github.com/01-edu/z01"

// Структура для хранения значений
type Values struct {
	FortyTwo  string
	TwentyOne string
}

// Структура точки
type point struct {
	x, y string
}

// Функция, изменяющая значения через структуру Values
func setPoint(ptr *point, v Values) {
	ptr.x = v.FortyTwo
	ptr.y = v.TwentyOne
}

func printStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	// Создаем структуру со строковыми значениями
	values := Values{
		FortyTwo:  "42",
		TwentyOne: "21",
	}

	points := &point{}
	setPoint(points, values)

	printStr("x = ")
	printStr(points.x)
	printStr(", y = ")
	printStr(points.y)
	z01.PrintRune('\n')
}
