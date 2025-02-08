package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0] // Start with the first string
	for i := 1; i < len(strs); i++ {
		result += sep + strs[i] // Add separator and next string
	}
	return result
}
