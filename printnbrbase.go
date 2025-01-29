package piscine

import "github.com/01-edu/z01"

// PrintNbrBase преобразует и выводит число в указанной системе счисления
func PrintNbrBase(nbr int, base string) {
	// Проверяем валидность базы
	if !isValidBase(base) {
		printString("NV")
		return
	}

	// Обрабатываем случай с отрицательным числом
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	// Вычисляем длину базы
	baseLen := len(base)

	// Если число равно нулю, выводим первый символ базы
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	// Преобразуем число в строку в указанной базе
	result := ""
	for nbr > 0 {
		remainder := nbr % baseLen
		result = string(base[remainder]) + result
		nbr /= baseLen
	}

	// Печатаем результат
	printString(result)
}

// Проверка валидности базы
func isValidBase(base string) bool {
	// База должна содержать минимум 2 символа
	if len(base) < 2 {
		return false
	}

	// Проверяем наличие дублирующихся символов или запрещённых символов (+ или -)
	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

// Вывод строки символ за символом
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
