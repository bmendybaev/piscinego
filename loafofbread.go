package piscine

func LoafOfBread(str string) string {
	filtered := ""
	for _, char := range str {
		if char != ' ' {
			filtered += string(char)
		}
	}

	if len(filtered) < 5 {
		return "Invalid Output\n"
	}

	result := ""

	for i := 0; i < len(filtered); {
		// Check if there are at least 5 characters left to process
		if i+5 <= len(filtered) {
			result += filtered[i:i+5] + " "
			i += 5 // Move forward by 5 characters
			if i < len(filtered) {
				i++ // Skip the next character if within bounds
			}
		} else {
			// Append the remaining characters
			result += filtered[i:]
			break
		}
	}

	// Remove the trailing space if present
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
