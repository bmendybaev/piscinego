package piscine

// ListRemoveIf removes all elements equal to data_ref from the list
func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil || l.Head == nil {
		return
	}

	// Remove nodes from the head if they match data_ref
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	// If the list becomes empty after removing head nodes
	if l.Head == nil {
		l.Tail = nil
		return
	}

	// Remove nodes beyond the head
	current := l.Head
	for current != nil && current.Next != nil {
		if current.Next.Data == data_ref {
			// Skip the node with data_ref
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	// Update the Tail reference
	l.Tail = l.Head
	for l.Tail != nil && l.Tail.Next != nil {
		l.Tail = l.Tail.Next
	}
}
