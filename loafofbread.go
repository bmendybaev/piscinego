package piscine

func LoafOfBread(str string) string {
	filtered := ""
	for _, char := range str {
		if char != ' ' {
			filtered += string(char)
		}
	}

	if len(filtered) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	word := ""
	count := 0

	for i := 0; i < len(filtered); i++ {
		word += string(filtered[i])
		count++

		if count == 5 {
			result += word + " "
			word = ""
			count = 0

			if i+1 < len(filtered) {
				i++ // Skip the next character
			}
		}
	}

	if word != "" {
		result += word
	}

	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
