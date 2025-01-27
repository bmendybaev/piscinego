package piscine

func IsUpper(s string) bool {
	// Iterate through each character in the string
	for i := 0; i < len(s); i++ {
		// Check if the character is not an uppercase letter
		if s[i] < 'A' || s[i] > 'Z' {
			return false
		}
	}
	// If all characters are uppercase, return true
	return true
}
