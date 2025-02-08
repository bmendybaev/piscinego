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

func normalizeString(str string) string {
	normalized := ""
	skipSpace := false

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			normalized += string(str[i])
			skipSpace = false
		} else if !skipSpace {
			normalized += " "
			skipSpace = true
		}
	}

	return normalized
}

func ShoppingSummaryCounter(str string) map[string]int {
	str = normalizeString(str)
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
