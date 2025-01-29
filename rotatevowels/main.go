package main

import (
	"os"

	"github.com/01-edu/z01"
)

// isVowel проверяет, является ли символ гласной (y не считается гласной)
func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

// rotateVowels зеркально меняет гласные по всей строке, игнорируя пробелы
func rotateVowels(s string) string {
	runes := []rune(s)
	vowels := []rune{}

	// Собираем все гласные
	for _, r := range runes {
		if isVowel(r) {
			vowels = append(vowels, r)
		}
	}

	// Переворачиваем порядок гласных
	vowelIndex := len(vowels) - 1
	for i := 0; i < len(runes); i++ {
		if isVowel(runes[i]) {
			runes[i] = vowels[vowelIndex]
			vowelIndex--
		}
	}

	return string(runes)
}

func main() {
	// Если нет аргументов, просто выводим новую строку
	if len(os.Args) <= 1 {
		z01.PrintRune('\n')
		return
	}

	// Объединяем все аргументы в одну строку с пробелами
	combined := ""
	for i, arg := range os.Args[1:] {
		combined += arg
		if i < len(os.Args[1:])-1 {
			combined += " "
		}
	}

	// Меняем гласные по всей строке (пробелы не учитываются)
	result := rotateVowels(combined)

	// Выводим результат
	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
