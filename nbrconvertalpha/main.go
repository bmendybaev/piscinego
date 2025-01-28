package main

import (
	"os"

	"github.com/01-edu/z01"
)

// customAtoi преобразует строку в целое число. Если строка некорректна, возвращает -1.
func customAtoi(s string) int {
	result := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return -1 // Если символ не является цифрой, возвращаем -1.
		}
		result = result*10 + int(r-'0') // Преобразуем руну в число.
	}
	return result
}

func main() {
	// Проверяем, используется ли флаг --upper
	upper := false
	start := 1
	if len(os.Args) > 1 && os.Args[1] == "--upper" {
		upper = true
		start = 2
	}

	// Проходим по аргументам, начиная с позиции start
	for i := start; i < len(os.Args); i++ {
		// Преобразуем аргумент в число с помощью customAtoi
		n := customAtoi(os.Args[i])
		if n < 1 || n > 26 {
			// Если число вне диапазона [1, 26] или аргумент некорректен, печатаем пробел
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
