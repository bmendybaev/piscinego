package piscine

import "unicode"

func capitalize(s string) string {
	result := ""
	for i, ch := range s {
		if i == 0 {
			result += string(unicode.ToUpper(ch))
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
