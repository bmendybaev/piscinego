package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	size := max - min
	result := make([]int, size) // Properly allocate the slice

	// Populate the slice
	for i := 0; i < size; i++ {
		result[i] = min + i
	}

	return result
}
