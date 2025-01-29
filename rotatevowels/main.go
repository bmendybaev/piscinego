package main

import (
	"fmt"
	"os"
	"strings"
)

func isVowel(c rune) bool {
	vowels := "AEIOUaeiou"
	return strings.ContainsRune(vowels, c)
}

func mirrorVowels(s string) string {
	runes := []rune(s)
	vowelIndices := []int{}
	vowelChars := []rune{}

	// Collect vowel indices and characters (ignoring spaces)
	for i, r := range runes {
		if isVowel(r) {
			vowelIndices = append(vowelIndices, i)
			vowelChars = append(vowelChars, r)
		}
	}

	// Reverse the vowels in place
	for i, j := 0, len(vowelChars)-1; i < j; i, j = i+1, j-1 {
		vowelChars[i], vowelChars[j] = vowelChars[j], vowelChars[i]
	}

	// Replace original vowels with mirrored ones
	vowelPos := 0
	for i := range runes {
		if isVowel(runes[i]) {
			runes[i] = vowelChars[vowelPos]
			vowelPos++
		}
	}

	return string(runes)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println()
		return
	}

	input := strings.Join(os.Args[1:], " ")
	fmt.Println(mirrorVowels(input))
}
