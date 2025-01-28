package main

import (
	"os"
	"sort"

	"github.com/01-edu/z01"
)

func main() {
	// Проверяем, есть ли аргументы
	if len(os.Args) > 1 {
		// Копируем аргументы в новый срез (исключая имя программы os.Args[0])
		args := os.Args[1:]

		// Сортируем аргументы в порядке ASCII
		sort.Strings(args)

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
