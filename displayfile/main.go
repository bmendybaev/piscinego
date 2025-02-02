package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Проверяем количество аргументов
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	// Открываем файл
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Cannot read file")
		return
	}
	defer file.Close()

	// Читаем и выводим содержимое файла
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println("Error reading file")
	}
}
