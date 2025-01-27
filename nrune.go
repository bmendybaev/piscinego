package piscine

func NRune(s string, n int) rune {
	// Return 0 if n is less than 1
	if n < 1 {
		return 0
	}
	// Iterate over the string to find the nth rune
	for i, r := range s {
		if i+1 == n { // i+1 because rune indices are 1-based
			return r
		}
	}
	// If the nth rune doesn't exist, return 0
	return 0
	// sdfdsfsdf
}
