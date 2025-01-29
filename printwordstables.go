package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, word := range a { // Remove unused `i`
		for _, char := range word {
			z01.PrintRune(char) // Print each character
		}
		z01.PrintRune('\n') // Print newline after each word
	}
}
