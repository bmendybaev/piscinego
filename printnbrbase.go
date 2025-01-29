package piscine

import "github.com/01-edu/z01"

// PrintNbrBase преобразует и выводит число в указанной системе счисления
func PrintNbrBase(nbr int, base string) {
	// Проверяем валидность базы
	if !isValidBase(base) {
		printString("NV")
		return
	}

	// Длина базы
	baseLen := len(base)

	// Обрабатываем случай с минимальным значением int64
	if nbr == -9223372036854775808 {
		z01.PrintRune('-')
		PrintNbrBase(int(-(nbr / baseLen)), base) // Рекурсивно обрабатываем часть числа
		z01.PrintRune(rune(base[-(nbr % baseLen)])) // Последняя цифра базы
		return
	}

	// Обрабатываем случай с отрицательным числом
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	// Если число равно нулю, выводим первый символ базы
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	// Преобразуем число в строку в указанной базе
	var result []rune
	n := nbr
	for n > 0 {
		remainder := n % baseLen
		result = append([]rune{rune(base[remainder])}, result...)
		n /= baseLen
	}

	// Печатаем результат
	for _, r := range result {
		z01.PrintRune(r)
	}
}

// Проверка валидности базы
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

// Вывод строки
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
