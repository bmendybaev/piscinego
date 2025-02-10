package piscine

// ListSize returns the number of elements in the linked list
func ListSize(l *List) int {
	count := 0
	current := l.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}
