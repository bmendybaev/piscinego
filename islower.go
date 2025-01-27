package piscine

func IsLower(s string) bool {
	// Iterate through each character in the string
	for i := 0; i < len(s); i++ {
		// Check if the character is not a lowercase letter
		if s[i] < 'a' || s[i] > 'z' {
			return false
		}
	}
	// If all characters are lowercase, return true
	return true
}
