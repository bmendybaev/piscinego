package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	size := 0
	for _, s := range args {
		size += len(s) + 1 // Account for the newline character
	}

	// Create a byte slice to efficiently build the string
	result := make([]byte, 0, size)

	for i, s := range args {
		result = append(result, s...)
		if i < len(args)-1 {
			result = append(result, '\n')
		}
	}

	return string(result)
}
