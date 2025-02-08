package piscine

func DescendAppendRange(max, min int) []int {
	// Initialize an empty slice to store the result
	result := []int{}

	// Check if max is greater than min
	if max > min {
		// Append values from max to min (excluding min)
		for i := max; i > min; i-- {
			result = append(result, i)
		}
	}

	// Return the resulting slice
	return result
}
