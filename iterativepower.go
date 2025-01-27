package piscine

func IterativePower(nb int, power int) int {
	// If power is negative, return 0
	if power < 0 {
		return 0
	}
	// If power is 0, return 1 (any number to the power of 0 is 1)
	if power == 0 {
		return 1
	}
	// Initialize result as 1
	result := 1
	// Multiply result by nb, power times
	for i := 0; i < power; i++ {
		result *= nb
	}
	return result
}
