package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	word := ""

	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if word != "" { // Ensure we don't add empty words
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}

	// Add the last word if there is one
	if word != "" {
		result = append(result, word)
	}

	return result
}
