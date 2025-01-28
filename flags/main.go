package main

import (
	"os"

	"github.com/01-edu/z01"
)

// sortString выполняет сортировку строки по ASCII
func sortString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		for j := 0; j < len(runes)-1-i; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}

// printHelp выводит инструкции по использованию программы
func printHelp() {
	helpText := `--insert
  -i
         This flag inserts the string into the string passed as argument.
--order
  -o
         This flag will behave like a boolean, if it is called it will order the argument.
`
	for _, r := range helpText {
		z01.PrintRune(r)
	}
}

func main() {
	// Если аргументов нет или указан флаг --help или -h
	if len(os.Args) == 1 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		printHelp()
		return
	}

	var insert string
	var order bool
	var mainString string

	// Обработка аргументов
	for _, arg := range os.Args[1:] {
		if len(arg) > 3 && (arg[:3] == "-i=" || arg[:9] == "--insert=") {
			// Обрабатываем вставку: извлекаем строку после = (для -i= или --insert=)
			equalIndex := 0
			for i, r := range arg {
				if r == '=' {
					equalIndex = i
					break
				}
			}
			insert = arg[equalIndex+1:]
		} else if arg == "-o" || arg == "--order" {
			order = true
		} else {
			// Всё остальное считается основной строкой
			mainString = arg
		}
	}

	// Если задан флаг --insert, добавляем строку
	if insert != "" {
		mainString += insert
	}

	// Если задан флаг --order, сортируем строку
	if order {
		mainString = sortString(mainString)
	}

	// Печатаем результат
	for _, r := range mainString {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
