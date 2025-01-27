package piscine

func Join(strs []string, sep string) string {
	// Handle empty slice
	if len(strs) == 0 {
		return ""
	}

	result := strs[0] // Start with the first string

	// Iterate through the rest of the slice and concatenate with the separator
	for i := 1; i < len(strs); i++ {
		result += sep + strs[i]
	}

	return result
}
