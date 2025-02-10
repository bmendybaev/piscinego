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

// ListLast returns the Data of the last element in the linked list
func ListLast(l *List) interface{} {
	if l.Tail != nil {
		return l.Tail.Data
	}
	return nil
}
