package main

import (
	"fmt"
	"os"
)

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
		fmt.Fprintf(os.Stderr, "open %s: no such file or directory\n", filename)
		return false
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "open %s: error retrieving file info\n", filename)
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
		fmt.Fprintln(os.Stderr, "error seeking file")
		return false
	}

	buffer := make([]byte, count)
	n, err := file.Read(buffer)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading file")
		return false
	}

	fmt.Print(string(buffer[:n]))
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 || args[0] != "-c" {
		fmt.Fprintln(os.Stderr, "Usage: go run . -c <bytes> <file1> [file2 ...]")
		os.Exit(1)
	}

	count, valid := atoi(args[1])
	if !valid || count <= 0 {
		fmt.Fprintln(os.Stderr, "Error: -c argument must be a positive number")
		os.Exit(1)
	}

	files := args[2:]
	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "Error: No files provided")
		os.Exit(1)
	}

	exitCode := 0
	firstFilePrinted := false
	errorPrinted := false

	for _, filename := range files {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			if errorPrinted {
				fmt.Fprintln(os.Stderr) // Перенос строки между ошибками
			}
			fmt.Fprintf(os.Stderr, "open %s: no such file or directory\n", filename)
			errorPrinted = true
			exitCode = 1
			continue
		}

		if firstFilePrinted || errorPrinted {
			fmt.Println() // Перенос строки перед заголовком файла
		}

		fmt.Printf("==> %s <==\n", filename)
		printTail(filename, count)

		firstFilePrinted = true
	}

	os.Exit(exitCode)
}
