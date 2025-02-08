package main

import (
	"github.com/01-edu/z01"
)

func rot14(input string) {
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			shifted := 'a' + (char-'a'+14)%26
			z01.PrintRune(shifted)
		} else if char >= 'A' && char <= 'Z' {
			shifted := 'A' + (char-'A'+14)%26
			z01.PrintRune(shifted)
		} else {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}

func main() {
	rot14("Hello, World!") // Output: Vszzc, Kcfzr!
}
