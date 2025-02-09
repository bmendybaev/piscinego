package piscine

func LoafOfBread(str string) string {
	// Remove all spaces from the input string
	filtered := ""
	for _, char := range str {
		if char != ' ' {
			filtered += string(char)
		}
	}

	// Check if the filtered string is less than 5 characters
	if len(filtered) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	index := 0

	// Process the filtered string in chunks of 5 characters
	for index+5 <= len(filtered) {
		result += filtered[index:index+5] + " "
		index += 5 // Move 5 characters forward
		if index < len(filtered) {
			index++ // Skip the next character if within bounds
		}
	}

	// Check for remaining characters if they form a valid chunk
	if len(filtered)-index >= 5 {
		result += filtered[index : index+5]
	}

	// Remove the trailing space if present and add a newline
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
