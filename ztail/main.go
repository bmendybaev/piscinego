package main

import (
	"fmt"
	"os"
)

// Функция для печати строки
func printStr(s string) {
	fmt.Print(s)
}

// Функция для преобразования строки в число (без strconv)
func atoi(s string) (int, bool) {
	result := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, false
		}
		result = result*10 + int(r-'0')
	}
	return result, true
}

// Функция для печати последних N байтов файла
func printTail(filename string, count int) bool {
	file, err := os.Open(filename)
	if err != nil {
		printStr("open " + filename + ": no such file or directory\n")
		return false
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		printStr("open " + filename + ": error retrieving file info\n")
		return false
	}

	if info.Size() == 0 {
		return true
	}

	start := info.Size() - int64(count)
	if start < 0 {
		start = 0
	}
	_, err = file.Seek(start, 0)
	if err != nil {
		printStr("error seeking file\n")
		return false
	}

	buffer := make([]byte, count)
	n, err := file.Read(buffer)
	if err != nil {
		printStr("error reading file\n")
		return false
	}

	fmt.Print(string(buffer[:n]))
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 || args[0] != "-c" {
		printStr("Usage: go run . -c <bytes> <file1> [file2 ...]\n")
		os.Exit(1)
	}

	count, valid := atoi(args[1])
	if !valid || count <= 0 {
		printStr("Error: -c argument must be a positive number\n")
		os.Exit(1)
	}

	files := args[2:]
	if len(files) == 0 {
		printStr("Error: No files provided\n")
		os.Exit(1)
	}

	exitCode := 0
	firstFilePrinted := false
	errorPrinted := false
	for _, filename := range files {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			if errorPrinted {
				printStr("\n") // Добавляем перенос строки между ошибками
			}
			printStr("open " + filename + ": no such file or directory")
			errorPrinted = true
			exitCode = 1
			continue
		}
		if firstFilePrinted || errorPrinted {
			printStr("\n") // Перенос строки перед печатью файла, если были ошибки
		}
		firstFilePrinted = true
		printStr("==> " + filename + " <==\n")
		if !printTail(filename, count) {
			exitCode = 1
		}
	}
	os.Exit(exitCode)
}
