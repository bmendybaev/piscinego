package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	start := 0

	for start < len(str) {
		// Skip any leading spaces
		for start < len(str) && str[start] == ' ' {
			start++
		}

		// Identify the end of the current word
		end := start
		for end < len(str) && str[end] != ' ' {
			end++
		}

		// Add the word to the map if it's not empty
		if start < end {
			word := str[start:end]
			summary[word]++
		}

		// Move to the next potential word and skip consecutive spaces
		start = end
		for start < len(str) && str[start] == ' ' {
			start++ // Skip consecutive spaces
		}
	}

	return summary
}
