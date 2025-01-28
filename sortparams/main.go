package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Проверяем, есть ли аргументы
	if len(os.Args) > 1 {
		// Копируем аргументы в новый срез (исключая имя программы os.Args[0])
		args := os.Args[1:]

		// Реализуем пузырьковую сортировку
		for i := 0; i < len(args)-1; i++ {
			for j := 0; j < len(args)-1-i; j++ {
				if args[j] > args[j+1] {
					// Меняем элементы местами
					args[j], args[j+1] = args[j+1], args[j]
				}
			}
		}

		// Выводим отсортированные аргументы
		for _, arg := range args {
			// Печатаем каждый символ аргумента
			for _, r := range arg {
				z01.PrintRune(r)
			}
			// Добавляем перевод строки после каждого аргумента
			z01.PrintRune('\n')
		}
	}
}
