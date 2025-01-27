package piscine

func LastRune(s string) rune {
	var lastRune rune
	// Iterate over the string and keep track of the last rune
	for _, r := range s {
		lastRune = r
	}
	// Return the last rune; if the string is empty, it returns the zero value (0)
	return lastRune
}
