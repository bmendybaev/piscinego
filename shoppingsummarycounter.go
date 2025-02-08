package piscine
func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	item := ""
	i := 0
	for i < len(str) {
		if str[i] != ' ' {
			item += string(str[i])
			i++
		} else {
			if item != "" {
				summary[item]++
				item = ""
			}
			// Skip all consecutive spaces
			for i < len(str) && str[i] == ' ' {
				i++
			}
		}
	}
	// Add the last item if there's no trailing space ghjgjhgjhgjhgj
	if item != "" {
		summary[item]++
	}
	return summary
}
