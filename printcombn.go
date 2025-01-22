package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	// Проверяем, что n находится в допустимых пределах
	if n <= 0 || n >= 10 {
		return
	}

	// Рекурсивная функция для генерации комбинаций
	var generate func(start, depth int, combination []rune)
	generate = func(start, depth int, combination []rune) {
		if depth == n {
			// Печатаем текущую комбинацию
			for _, r := range combination {
				z01.PrintRune(r)
			}
			// Проверяем, последняя ли это комбинация
			if combination[0] != '9'-rune(n)+1 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			return
		}

		// Рекурсивно добавляем цифры в комбинацию
		for i := start; i <= 9; i++ {
			generate(i+1, depth+1, append(combination, '0'+rune(i)))
		}
	}

	// Запускаем генерацию комбинаций
	generate(0, 0, []rune{})
	// Завершаем строку
	z01.PrintRune('\n')
}
