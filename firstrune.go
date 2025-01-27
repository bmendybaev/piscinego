package piscine

func FirstRune(s string) rune {
	// Iterate over the string and return the first rune
	for _, r := range s {
		return r
	}
	// If the string is empty, return 0 (default rune value)
	return 0
}
