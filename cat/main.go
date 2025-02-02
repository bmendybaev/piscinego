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

// Читает и выводит содержимое файла
func printFile(filename string) bool {
	file, err := os.Open(filename)
	if err != nil {
		printString("ERROR: " + err.Error() + "\n")
		return false
	}
	defer file.Close()

	// Копируем содержимое файла в stdout
	io.Copy(os.Stdout, file)
	return true
}

// Читает stdin и печатает ввод
func readStdin() {
	io.Copy(os.Stdout, os.Stdin)
}

func main() {
	args := os.Args[1:]

	// Если аргументов нет, читаем stdin
	if len(args) == 0 {
		readStdin()
		return
	}

	exitCode := 0 // Код выхода (0 — успех, 1 — ошибка)

	// Читаем файлы по очереди
	for _, filename := range args {
		if !printFile(filename) {
			exitCode = 1
		}
	}

	// Если были ошибки, завершаем программу с кодом 1
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}
