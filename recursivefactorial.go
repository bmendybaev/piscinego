package piscine

func RecursiveFactorial(nb int) int {
	// Return 0 for invalid inputs or values that can cause overflow
	if nb < 0 || nb > 20 {
		return 0
	}
	// Base case: factorial of 0 or 1 is 1
	if nb == 0 || nb == 1 {
		return 1
	}
	// Recursive case
	return nb * RecursiveFactorial(nb-1)
}
