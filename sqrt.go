package piscine

func Sqrt(nb int) int {
	// Return 0 for negative numbers
	if nb < 0 {
		return 0
	}
	// Iterate from 1 up to nb (or until the square root is found)
	for i := 1; i*i <= nb; i++ {
		// If i*i equals nb, i is the square root
		if i*i == nb {
			return i
		}
	}
	// If no whole number square root exists, return 0
	return 0
}
