package main

import (
	"fmt"
)

// Структура точки
type point struct {
	x, y int
}

// Функция для установки значений (можно менять x и y)
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points) // Можно менять points.x и points.y

	// Вывод результата через fmt.Printf
	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
