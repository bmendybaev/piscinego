package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	size := max - min
	result := []int{} // Create an empty slice

	// Extend the slice to the required size
	result = result[:size]

	// Populate the slice
	for i := range result {
		result[i] = min + i
	}

	return result
}
