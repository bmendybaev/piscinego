package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	item := ""
	i := 0

	for i < len(str) {
		// Skip leading spaces
		for i < len(str) && str[i] == ' ' {
			i++
		}

		// Build the current item
		for i < len(str) && str[i] != ' ' {
			item += string(str[i])
			i++
		}

		if item != "" {
			summary[item]++
			item = "" // Reset for the next word 123
		}
	}

	return summary
}
