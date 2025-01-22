package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	// Задаём цифры вручную как массив символов
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	// Перебираем первую пару цифр
	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(digits); j++ {
			// Перебираем вторую пару цифр
			for k := i; k < len(digits); k++ {
				for l := 0; l < len(digits); l++ {
					// Проверяем, чтобы первая пара была меньше второй
					if i < k || (i == k && j < l) {
						// Печатаем первую пару
						z01.PrintRune(digits[i])
						z01.PrintRune(digits[j])
						// Печатаем пробел
						z01.PrintRune(' ')
						// Печатаем вторую пару
						z01.PrintRune(digits[k])
						z01.PrintRune(digits[l])

						// Добавляем запятую и пробел, кроме последней комбинации
						if i != 9 || j != 8 || k != 9 || l != 9 {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	// Перевод строки в конце
	z01.PrintRune('\n')
}
