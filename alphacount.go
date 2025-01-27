package piscine

func AlphaCount(s string) int {
	count := 0
	// Iterate through each character in the string
	for _, r := range s {
		// Check if the rune is a letter (Latin alphabet)
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			count++
		}
	}
	return count
}
