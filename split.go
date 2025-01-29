package piscine

func Split(s, sep string) []string {
	if sep == "" {
		return []string{s} // If separator is empty, return the whole string as a single element
	}

	var result []string
	sepLen := len(sep)
	start := 0

	// Determine the maximum number of substrings
	count := 1
	for i := 0; i <= len(s)-sepLen; i++ {
		if s[i:i+sepLen] == sep {
			count++
			i += sepLen - 1 // Move past the separator
		}
	}

	// Allocate the slice using `make`
	result = make([]string, 0, count)

	// Perform the split operation
	for i := 0; i <= len(s)-sepLen; i++ {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i]) // Append substring
			start = i + sepLen                  // Move past the separator
			i += sepLen - 1                     // Move iterator forward
		}
	}

	// Append the last part of the string
	result = append(result, s[start:])

	return result
}
