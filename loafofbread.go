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
	i := 0
	for i+5 <= len(filtered) {
		result += filtered[i:i+5] + " "
		i += 5
		if i < len(filtered) {
			i++ // Skip the next character
		}
	}

	// Append remaining characters if any
	if i < len(filtered) {
		remaining := filtered[i:]
		if len(remaining) >= 5 {
			result += remaining[:5]
		} else {
			result += remaining
		}
	}

	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
