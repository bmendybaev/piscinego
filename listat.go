package piscine

// ListAt returns the pointer to the NodeL at the given position in the linked list
func ListAt(l *NodeL, pos int) *NodeL {
	current := l
	index := 0

	for current != nil {
		if index == pos {
			return current
		}
		current = current.Next
		index++
	}

	return nil
}
