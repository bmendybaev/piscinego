package piscine

// Compare function compares two strings in lexicographical order.
func Compare(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
