package piscine
import main

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 { // Handle invalid values or overflow
		return 0
	}

	result := 1
	for i := 1; i <= nb; i++ {
		result *= i
	}

	return result
}
