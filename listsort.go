package piscine

// NodeI represents a node in the linked list with integer data
type NodeI struct {
	Data int
	Next *NodeI
}

// ListSort sorts the linked list in ascending order
func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	swapped := true

	for swapped {
		swapped = false
		current := l

		for current.Next != nil {
			if current.Data > current.Next.Data {
				// Swap the data between current and next nodes
				current.Data, current.Next.Data = current.Next.Data, current.Data
				swapped = true
			}
			current = current.Next
		}
	}

	return l
}
