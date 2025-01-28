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

// reverseVowels зеркально меняет гласные во всей строке
func reverseVowels(s string) string {
	runes := []rune(s) // Преобразуем строку в срез рун
	left, right := 0, len(runes)-1

	for left < right {
		// Найти первую гласную слева
		for left < right && !isVowel(runes[left]) {
			left++
		}
		// Найти первую гласную справа
		for left < right && !isVowel(runes[right]) {
			right--
		}
		// Поменять местами гласные
		if left < right {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
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

	// Зеркально меняем гласные во всей строке
	result := reverseVowels(combined)

	// Выводим результат
	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
