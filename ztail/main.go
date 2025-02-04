package main

import (
	"os"
	"github.com/01-edu/z01"
)

// Функция для печати строки
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

// Функция для печати числа
func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	var rev [10]rune
	idx := 0
	for n > 0 {
		rev[idx] = '0' + rune(n%10)
		n /= 10
		idx++
	}
	for i := idx - 1; i >= 0; i-- {
		z01.PrintRune(rev[i])
	}
}

// Функция для преобразования строки в число (без strconv)
func atoi(s string) (int, bool) {
	result := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, false // Ошибка: не цифра
		}
		result = result*10 + int(r-'0')
	}
	return result, true
}

// Функция для печати последних N байтов файла
func printTail(filename string, count int) bool {
	file, err := os.Open(filename)
	if err != nil {
		printStr("open ")
		printStr(filename)
		printStr(": no such file or directory\n")
		return false
	}
	defer file.Close()

	// Получаем размер файла
	info, err := file.Stat()
	if err != nil {
		printStr("open ")
		printStr(filename)
		printStr(": error retrieving file info\n")
		return false
	}

	start := info.Size() - int64(count)
	if start < 0 {
		start = 0
	}
	file.Seek(start, 0)
	buffer := make([]byte, count)
	n, _ := file.Read(buffer)
	for i := 0; i < n; i++ {
		z01.PrintRune(rune(buffer[i]))
	}
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

	firstFilePrinted := false
	for _, filename := range files {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			printStr("open ")
			printStr(filename)
			printStr(": no such file or directory\n")
			continue
		}
		if firstFilePrinted {
			z01.PrintRune('\n')
		}
		firstFilePrinted = true
		printStr("==> ")
		printStr(filename)
		printStr(" <==\n")
		printTail(filename, count)
	}
}