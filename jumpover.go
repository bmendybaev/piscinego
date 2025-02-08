package piscine

import "github.com/01-edu/z01"

func JumpOver(str string) string {
	result := ""

	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	// Отладочный вывод для проверки индексов и символов
	for i := 2; i < len(str); i += 3 {
		z01.PrintRune(rune(str[i]))
	}
	z01.PrintRune('\n')

	return result + "\n"
}
