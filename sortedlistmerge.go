package piscine

// SortedListMerge merges two sorted linked lists into one sorted linked list
func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	// Create a dummy node to simplify edge cases
	dummy := &NodeI{}
	current := dummy

	// Merge the two lists
	for n1 != nil && n2 != nil {
		if n1.Data < n2.Data {
			current.Next = n1
			n1 = n1.Next
		} else {
			current.Next = n2
			n2 = n2.Next
		}
		current = current.Next
	}

	// Attach any remaining nodes
	if n1 != nil {
		current.Next = n1
	} else {
		current.Next = n2
	}

	// Return the merged list (skipping the dummy node)
	return dummy.Next
}
