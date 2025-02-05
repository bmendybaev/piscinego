package piscine

// IsSorted checks if the slice a is sorted in ascending or descending order according to the function f
func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := true
	descending := true

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			ascending = false
		}
		if f(a[i], a[i+1]) < 0 {
			descending = false
		}
	}
	return ascending || descending
}

// IsSortedBy10 compares numbers based on their last digit (modulo 10)
func IsSortedBy10(a, b int) int {
	return (a % 10) - (b % 10)
}
