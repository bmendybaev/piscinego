package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	start := -1 // Marker for the start of a word

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			if start == -1 {
				start = i // Mark the beginning of a word
			}
		}

		// If a space is encountered or it's the end of the string, process the word
		if (str[i] == ' ' && start != -1) || (i == len(str)-1 && start != -1) {
			end := i
			if str[i] != ' ' {
				end = i + 1 // Include the last character if not a space
			}
			word := str[start:end]
			summary[word]++
			start = -1 // Reset for the next word
		}
	}

	return summary
}
