package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

// Функция для вывода строки через z01.PrintRune
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

// Читает и выводит содержимое файла
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		printString("ERROR: " + err.Error() + "\n")
		return
	}
	defer file.Close()

	// Копируем файл в stdout без bufio
	io.Copy(os.Stdout, file)
}

// Читает stdin и выводит обратно
func readStdin() {
	io.Copy(os.Stdout, os.Stdin)
}

func main() {
	args := os.Args[1:]

	// Если аргументов нет → читаем из stdin
	if len(args) == 0 {
		readStdin()
		return
	}

	// Читаем файлы по очереди
	for _, filename := range args {
		printFile(filename)
	}
}
