package piscine

func ToLower(s string) string {
	// Convert the string to a mutable slice of bytes
	result := []byte(s)

	// Iterate through each character in the string
	for i := 0; i < len(result); i++ {
		// Check if the character is an uppercase letter
		if result[i] >= 'A' && result[i] <= 'Z' {
			// Convert it to lowercase by adding 32 (ASCII conversion)
			result[i] += 32
		}
	}

	// Convert the byte slice back to a string and return it
	return string(result)
}
