package piscine

func RecursivePower(nb int, power int) int {
	// If power is negative, return 0
	if power < 0 {
		return 0
	}
	// Base case: Any number to the power of 0 is 1
	if power == 0 {
		return 1
	}
	// Recursive case
	return nb * RecursivePower(nb, power-1)
}
