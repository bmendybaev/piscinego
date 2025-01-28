package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	// Проверяем, передан ли флаг --upper
	upper := false
	start := 1
	if len(os.Args) > 1 && os.Args[1] == "--upper" {
		upper = true
		start = 2
	}

	// Проходим по аргументам, начиная с позиции start
	for i := start; i < len(os.Args); i++ {
		// Пробуем преобразовать аргумент в число
		n, err := strconv.Atoi(os.Args[i])
		if err != nil || n < 1 || n > 26 {
			// Если ошибка или число вне диапазона, печатаем пробел
			z01.PrintRune(' ')
		} else {
			// Преобразуем позицию в букву
			var letter rune
			if upper {
				letter = 'A' + rune(n-1) // Верхний регистр
			} else {
				letter = 'a' + rune(n-1) // Нижний регистр
			}
			z01.PrintRune(letter)
		}
	}

	// Перевод строки в конце
	z01.PrintRune('\n')
}
