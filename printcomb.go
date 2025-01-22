package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for k := j + 1; k <= '9'; k++ {
				// Печатаем три цифры
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)

				// Добавляем запятую и пробел, кроме последней комбинации
				if i != '7' || j != '8' || k != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	// Перевод строки в конце
	z01.PrintRune('\n')
}