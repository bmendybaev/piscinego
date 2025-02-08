package piscine

func capitalize(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if i == 0 && ch >= 'a' && ch <= 'z' {
			result += string(ch - ('a' - 'A'))
		} else {
			result += string(ch)
		}
	}
	return result
}

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	mergedWord := ""
	spaceCount := 0

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			if spaceCount > 0 && mergedWord != "" {
				capitalizedWord := capitalize(mergedWord)
				summary[capitalizedWord]++
				mergedWord = ""
			}
			mergedWord += string(str[i])
			spaceCount = 0
		} else {
			spaceCount++
		}
	}

	// Add the last merged word if there's no trailing space
	if mergedWord != "" {
		capitalizedWord := capitalize(mergedWord)
		summary[capitalizedWord]++
	}

	return summary
}
