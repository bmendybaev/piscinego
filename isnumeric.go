package piscine

func IsNumeric(s string) bool {
	// Iterate through each character in the string by index
	for i := 0; i < len(s); i++ {
		// Check if the character is not a digit
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	// If all characters are digits or the string is empty, return true
	return true
}
