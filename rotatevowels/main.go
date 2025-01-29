package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(c rune) bool {
	vowels := "AEIOUaeiou"
	for _, v := range vowels {
		if v == c {
			return true
		}
	}
	return false
}

func mirrorVowels(s []rune) {
	var vowels []rune
	for _, r := range s {
		if isVowel(r) {
			vowels = append(vowels, r)
		}
	}

	// Reverse collected vowels
	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}

	// Replace vowels in original string
	vowelPos := 0
	for i, r := range s {
		if isVowel(r) {
			s[i] = vowels[vowelPos]
			vowelPos++
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	}

	// Concatenate all arguments with spaces
	var input []rune
	for i, arg := range os.Args[1:] {
		if i > 0 {
			input = append(input, ' ')
		}
		input = append(input, []rune(arg)...)
	}

	mirrorVowels(input)

	for _, r := range input {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
// 
}
