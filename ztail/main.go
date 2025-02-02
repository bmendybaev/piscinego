package main

import (
	"fmt"
	"os"
	"strconv"
)

// Функция для печати последних N байтов файла
func printTail(filename string, count int) bool {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open %s: %v\n", filename, err)
		return false
	}
	defer file.Close()

	// Получаем размер файла
	info, err := file.Stat()
	if err != nil {
		fmt.Printf("open %s: %v\n", filename, err)
		return false
	}

	// Если count больше размера файла, читаем весь файл
	start := info.Size() - int64(count)
	if start < 0 {
		start = 0
	}

	// Считываем последние count байтов
	file.Seek(start, 0)
	buffer := make([]byte, count)
	n, _ := file.Read(buffer)

	// Выводим содержимое файла (без лишних переносов строки)
	fmt.Print(string(buffer[:n]))
	return true
}

func main() {
	args := os.Args[1:]

	// Проверка аргументов
	if len(args) < 2 || args[0] != "-c" {
		fmt.Println("Usage: go run . -c <bytes> <file1> [file2 ...]")
		os.Exit(1)
	}

	// Парсим количество байтов
	count, err := strconv.Atoi(args[1])
	if err != nil || count <= 0 {
		fmt.Println("Error: -c argument must be a positive number")
		os.Exit(1)
	}

	files := args[2:]
	if len(files) == 0 {
		fmt.Println("Error: No files provided")
		os.Exit(1)
	}

	firstFilePrinted := false

	// Обрабатываем все файлы
	for _, filename := range files {
		// Если файл не существует, просто печатаем ошибку и продолжаем
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			fmt.Printf("open %s: no such file or directory\n", filename)
			continue
		}

		// Разделяем файлы новой строкой, кроме первого файла
		if firstFilePrinted {
			fmt.Print("\n")
		}
		firstFilePrinted = true

		// Вывод заголовка файла
		fmt.Printf("==> %s <==\n", filename)
		printTail(filename, count)
	}
}
