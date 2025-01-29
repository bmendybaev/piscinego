package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i, word := range a {
		for _, char := range word {
			z01.PrintRune(char) // Print each character
		}
		z01.PrintRune('\n') // Print newline after each word
	}
}
