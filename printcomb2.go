package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 99; i++ {
		for j := i + 1; j <= 99; j++ {
			// Вывод первой цифры первого числа
			z01.PrintRune('0' + rune(i/10))
			// Вывод второй цифры первого числа
			z01.PrintRune('0' + rune(i%10))
			// Пробел между числами
			z01.PrintRune(' ')
			// Вывод первой цифры второго числа
			z01.PrintRune('0' + rune(j/10))
			// Вывод второй цифры второго числа
			z01.PrintRune('0' + rune(j%10))

			// Добавляем запятую и пробел, кроме последней комбинации
			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	// Перевод строки в конце
	z01.PrintRune('\n')
}
