package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Проверяем, есть ли аргументы (os.Args[1:])
	if len(os.Args) > 1 {
		// Перебираем аргументы в обратном порядке
		for i := len(os.Args) - 1; i > 0; i-- {
			// Печатаем каждый символ аргумента
			for _, r := range os.Args[i] {
				z01.PrintRune(r)
			}
			// Добавляем перевод строки после каждого аргумента
			z01.PrintRune('\n')
		}
	}
}
