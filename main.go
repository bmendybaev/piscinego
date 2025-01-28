package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Проверяем, есть ли аргументы (os.Args[1:])
	if len(os.Args) > 1 {
		// Перебираем каждый аргумент, начиная с первого (ос.Args[1:])
		for _, arg := range os.Args[1:] {
			// Печатаем каждый символ аргумента с помощью z01.PrintRune
			for _, r := range arg {
				z01.PrintRune(r)
			}
			// После каждого аргумента добавляем перевод строки
			z01.PrintRune('\n')
		}
	}
}
