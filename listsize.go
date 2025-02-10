package piscine

// NodeL represents a node in the linked list
type NodeL struct {
	Data interface{}
	Next *NodeL
}

// List represents the linked list
type List struct {
	Head *NodeL
	Tail *NodeL
}

// ListPushFront adds a new element to the front of the list
func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
}

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
