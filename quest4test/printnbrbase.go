package piscine

import "github.com/01-edu/z01"

// PrintNbrBase преобразует и выводит число в заданной системе счисления
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
	baseLen := 0
	for range base {
		baseLen++
	}

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
	baseLen := 0
	for range base {
		baseLen++
	}
	if baseLen < 2 {
		return false
	}

	// Проверяем наличие дублирующихся символов или запрещённых символов (+ или -)
	for i, r1 := range base {
		if r1 == '+' || r1 == '-' {
			return false
		}
		for j, r2 := range base {
			if i != j && r1 == r2 {
				return false
			}
		}
	}

	return true
}

// Вывод строки символ за символом
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
