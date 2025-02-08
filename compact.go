package piscine

func Compact(ptr *[]string) int {
	// Create a new slice to hold non-zero (non-empty) values
	nonZeroValues := []string{}

	// Iterate over the original slice
	for _, value := range *ptr {
		if value != "" {
			nonZeroValues = append(nonZeroValues, value)
		}
	}

	// Replace the original slice with the compacted version
	*ptr = nonZeroValues

	// Return the count of non-zero elements
	return len(nonZeroValues)
}
