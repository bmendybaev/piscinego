package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	item := ""

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			item += string(str[i])
		} else {
			if item != "" {
				summary[item]++
				item = ""
			}
			// Skip consecutive spaces without adding empty items
			for i+1 < len(str) && str[i+1] == ' ' {
				i++
			}
		}
	}

	// Add the last item if there's no trailing space
	if item != "" {
		summary[item]++
	}

	return summary
}
