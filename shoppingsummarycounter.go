package piscine

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	if s[0] >= 'a' && s[0] <= 'z' {
		return string(s[0]-('a'-'A')) + s[1:]
	}
	return s
}

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""

	for i := 0; i <= len(str); i++ {
		if i < len(str) && str[i] != ' ' {
			word += string(str[i])
		} else {
			if word != "" {
				capitalizedWord := capitalize(word)
				summary[capitalizedWord]++
				word = ""
			}
			// Skip multiple spaces
			for i+1 < len(str) && str[i+1] == ' ' {
				i++
			}
		}
	}

	return summary
}
