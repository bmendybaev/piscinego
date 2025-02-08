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

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			mergedWord += string(str[i])
		} else if mergedWord != "" {
			capitalizedWord := capitalize(mergedWord)
			summary[capitalizedWord]++
			mergedWord = ""
		}
	}

	// Add the last merged word if there's no trailing space
	if mergedWord != "" {
		capitalizedWord := capitalize(mergedWord)
		summary[capitalizedWord]++
	}

	return summary
}
