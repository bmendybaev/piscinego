package piscine

// CompStr compares two interface{} values for equality
func CompStr(a, b interface{}) bool {
	return a == b
}

// ListFind returns the address of the data of the first node that matches ref
func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current := l.Head
	for current != nil {
		if comp(current.Data, ref) {
			return &current.Data
		}
		current = current.Next
	}
	return nil
}
