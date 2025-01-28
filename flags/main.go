package main

import (
	"os"

	"github.com/01-edu/z01"
)

// sortString выполняет сортировку строки по ASCII.
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

// printHelp выводит инструкции по использованию программы.
func printHelp() {
	z01.PrintRune('-')
	z01.PrintRune('-')
	z01.PrintRune('i')
	z01.PrintRune('n')
	z01.PrintRune('s')
	z01.PrintRune('e')
	z01.PrintRune('r')
	z01.PrintRune('t')
	z01.PrintRune('\n')
	z01.PrintRune(' ')
	z01.PrintRune(' ')
	z01.PrintRune('-')
	z01.PrintRune('i')
	z01.PrintRune('\n')
	z01.PrintRune('\t')
	z01.PrintRune(' ')
	z01.PrintRune('T')
	z01.PrintRune('h')
	z01.PrintRune('i')
	z01.PrintRune('s')
	z01.PrintRune(' ')
	z01.PrintRune('f')
	z01.PrintRune('l')
	z01.PrintRune('a')
	z01.PrintRune('g')
	z01.PrintRune(' ')
	z01.PrintRune('i')
	z01.PrintRune('n')
	z01.PrintRune('s')
	z01.PrintRune('e')
	z01.PrintRune('r')
	z01.PrintRune('t')
	z01.PrintRune('s')
	z01.PrintRune(' ')
	z01.PrintRune('t')
	z01.PrintRune('h')
	z01.PrintRune('e')
	z01.PrintRune(' ')
	z01.PrintRune('s')
	z01.PrintRune('t')
	z01.PrintRune('r')
	z01.PrintRune('i')
	z01.PrintRune('n')
	z01.PrintRune('g')
	z01.PrintRune(' ')
	z01.PrintRune('i')
	z01.PrintRune('n')
	z01.PrintRune('t')
	z01.PrintRune('o')
	z01.PrintRune(' ')
	z01.PrintRune('t')
	z01.PrintRune('h')
	z01.PrintRune('e')
	z01.PrintRune(' ')
	z01.PrintRune('s')
	z01.PrintRune('t')
	z01.PrintRune('r')
	z01.PrintRune('i')
	z01.PrintRune('n')
	z01.PrintRune('g')
	z01.PrintRune(' ')
	z01.PrintRune('p')
	z01.PrintRune('a')
	z01.PrintRune('s')
	z01.PrintRune('s')
	z01.PrintRune('e')
	z01.PrintRune('d')
	z01.PrintRune(' ')
	z01.PrintRune('a')
	z01.PrintRune('s')
	z01.PrintRune(' ')
	z01.PrintRune('a')
	z01.PrintRune('r')
	z01.PrintRune('g')
	z01.PrintRune('u')
	z01.PrintRune('m')
	z01.PrintRune('e')
	z01.PrintRune('n')
	z01.PrintRune('t')
	z01.PrintRune('.')
	z01.PrintRune('\n')

	z01.PrintRune('-')
	z01.PrintRune('-')
	z01.PrintRune('o')
	z01.PrintRune('r')
	z01.PrintRune('d')
	z01.PrintRune('e')
	z01.PrintRune('r')
	z01.PrintRune('\n')
	z01.PrintRune(' ')
	z01.PrintRune(' ')
	z01.PrintRune('-')
	z01.PrintRune('o')
	z01.PrintRune('\n')
	z01.PrintRune('\t')
	z01.PrintRune(' ')
	z01.PrintRune('T')
	z01.PrintRune('h')
	z01.PrintRune('i')
	z01.PrintRune('s')
	z01.PrintRune(' ')
	z01.PrintRune('f')
	z01.PrintRune('l')
	z01.PrintRune('a')
	z01.PrintRune('g')
	z01.PrintRune(' ')
	z01.PrintRune('w')
	z01.PrintRune('i')
	z01.PrintRune('l')
	z01.PrintRune('l')
	z01.PrintRune(' ')
	z01.PrintRune('b')
	z01.PrintRune('e')
	z01.PrintRune('h')
	z01.PrintRune('a')
	z01.PrintRune('v')
	z01.PrintRune('e')
	z01.PrintRune(' ')
	z01.PrintRune('l')
	z01.PrintRune('i')
	z01.PrintRune('k')
	z01.PrintRune('e')
	z01.PrintRune(' ')
	z01.PrintRune('a')
	z01.PrintRune(' ')
	z01.PrintRune('b')
	z01.PrintRune('o')
	z01.PrintRune('o')
	z01.PrintRune('l')
	z01.PrintRune('e')
	z01.PrintRune('a')
	z01.PrintRune('n')
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('i')
	z01.PrintRune('f')
	z01.PrintRune(' ')
	z01.PrintRune('i')
	z01.PrintRune('t')
	z01.PrintRune(' ')
	z01.PrintRune('i')
	z01.PrintRune('s')
	z01.PrintRune(' ')
	z01.PrintRune('c')
	z01.PrintRune('a')
	z01.PrintRune('l')
	z01.PrintRune('l')
	z01.PrintRune('e')
	z01.PrintRune('d')
	z01.PrintRune(' ')
	z01.PrintRune('i')
	z01.PrintRune('t')
	z01.PrintRune(' ')
	z01.PrintRune('w')
	z01.PrintRune('i')
	z01.PrintRune('l')
	z01.PrintRune('l')
	z01.PrintRune(' ')
	z01.PrintRune('o')
	z01.PrintRune('r')
	z01.PrintRune('d')
	z01.PrintRune('e')
	z01.PrintRune('r')
	z01.PrintRune(' ')
	z01.PrintRune('t')
	z01.PrintRune('h')
	z01.PrintRune('e')
	z01.PrintRune(' ')
	z01.PrintRune('a')
	z01.PrintRune('r')
	z01.PrintRune('g')
	z01.PrintRune('u')
	z01.PrintRune('m')
	z01.PrintRune('e')
	z01.PrintRune('n')
	z01.PrintRune('t')
	z01.PrintRune('.')
	z01.PrintRune('\n')
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
		if len(arg) > 9 && arg[:9] == "--insert=" {
			insert = arg[9:]
		} else if arg == "-i" || arg == "--insert" {
			insert = ""
		} else if arg == "-o" || arg == "--order" {
			order = true
		} else {
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
