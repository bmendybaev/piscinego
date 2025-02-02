package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

// Функция для печати строки через z01.PrintRune
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

// Читает и выводит содержимое файла (выходит с ошибкой, если файл не найден)
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		printString("ERROR: " + err.Error() + "\n")
		os.Exit(1) // Если файла нет, сразу выходим с exit status 1
	}
	defer file.Close()

	// Читаем файл побайтово и печатаем
	io.Copy(os.Stdout, file)
}

func main() {
	args := os.Args[1:]

	// Если аргументов нет, читаем stdin
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	// Читаем файлы по очереди
	for _, filename := range args {
		printFile(filename)
	}
}
