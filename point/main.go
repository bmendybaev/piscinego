package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func toString(n int) string {
	if n == 0 {
		return "0"
	}
	str := ""
	if n < 0 {
		str += "-"
		n = -n
	}
	for temp := n; temp > 0; temp /= 10 {
		str = string(rune('0'+temp%10)) + str
	}
	return str
}

func main() {
	points := &point{}
	setPoint(points)

	// Формируем строку перед выводом
	output := "x = " + toString(points.x) + ", y = " + toString(points.y) + "\n"

	// Вывод ≤ 9 вызовов PrintRune
	for _, r := range output {
		z01.PrintRune(r)
	}
}
