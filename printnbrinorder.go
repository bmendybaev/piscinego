package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	// Handle the special case where n is 0
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Create a slice to store the digits
	digits := []int{}

	// Extract digits from the number
	for n > 0 {
		digits = append(digits, n%10) // Add the last digit to the slice
		n /= 10                       // Remove the last digit
	}

	// Sort the digits in ascending order
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			if digits[i] > digits[j] {
				digits[i], digits[j] = digits[j], digits[i] // Swap digits
			}
		}
	}

	// Print the digits
	for _, digit := range digits {
		z01.PrintRune(rune('0' + digit))
	}
}
