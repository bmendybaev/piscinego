package piscine

func IsPrime(nb int) bool {
	// 1 is not a prime number, and negative numbers are not considered prime
	if nb <= 1 {
		return false
	}
	// 2 and 3 are prime numbers
	if nb == 2 || nb == 3 {
		return true
	}
	// Exclude even numbers and multiples of 3
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}
	// Check divisors from 5 to sqrt(nb), skipping even numbers
	for i := 5; i*i <= nb; i += 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
	}
	return true
}
