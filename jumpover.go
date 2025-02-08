package piscine

import "github.com/01-edu/z01"

func JumpOver(str string) string {
	// Проверка на пустую строку и ограничения
	if len(str) < 3 {
		z01.PrintRune('\n')
		return "\n"
	}

	// Разбор строки и сборка заново
	runes := []rune(str)
	result := ""

	// Отбор каждого третьего символа
	for i := 2; i < len(runes); i += 3 {
		result += string(runes[i])
	}

	// Проверка на наличие выбранных символов
	if result == "" {
		z01.PrintRune('\n')
		return "\n"
	}

	// Вывод результата
	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')

	return result + "\n"
}
