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
	for i := 0; i+5 <= len(filtered); i += 6 {
		result += filtered[i:i+5] + " "
	}

	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
