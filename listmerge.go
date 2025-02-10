package piscine

// ListMerge appends elements of list l2 to the end of list l1
func ListMerge(l1 *List, l2 *List) {
	if l1 == nil || l2 == nil || l2.Head == nil {
		return
	}

	// If l1 is empty, set l1's head and tail to l2's head and tail
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}

	// Merge l2 at the end of l1
	l1.Tail.Next = l2.Head
	l1.Tail = l2.Tail
}
