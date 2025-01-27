package piscine

func IsPrintable(s string) bool {
	// Iterate through each character in the string
	for i := 0; i < len(s); i++ {
		// Check if the character is not a printable ASCII character
		if s[i] < 32 || s[i] > 126 {
			return false
		}
	}
	// If all characters are printable, return true
	return true
}
