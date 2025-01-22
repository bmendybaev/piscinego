package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 99; i++ {
		for j := i + 1; j <= 99; j++ {
			// Печатаем первое двузначное число
			z01.PrintRune(rune(i/10 + '0')) // Первая цифра
			z01.PrintRune(rune(i%10 + '0')) // Вторая цифра
			z01.PrintRune(' ')              // Пробел между числами

			// Печатаем второе двузначное число
			z01.PrintRune(rune(j/10 + '0')) // Первая цифра
			z01.PrintRune(rune(j%10 + '0')) // Вторая цифра

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
