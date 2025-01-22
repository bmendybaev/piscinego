ppackage piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	// Обрабатываем случай минимального значения int
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		PrintNbr(922337203685477580)
		z01.PrintRune('8')
		return
	}

	// Обрабатываем отрицательные числа
	if n < 0 {
		z01.PrintRune('-')
		n = -n // Преобразуем в положительное
	}

	// Рекурсивно печатаем цифры
	if n >= 10 {
		PrintNbr(n / 10) // Печатаем старшие цифры
	}
	z01.PrintRune('0' + rune(n%10)) // Печатаем младшую цифру
}
