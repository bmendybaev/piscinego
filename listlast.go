package piscine

// ListLast returns the Data of the last element in the linked list
func ListLast(l *List) interface{} {
	if l.Tail != nil {
		return l.Tail.Data
	}
	return nil
}
