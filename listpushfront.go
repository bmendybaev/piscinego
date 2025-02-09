package piscine

// NodeL represents a node in the linked list
type NodeL struct {
	Data interface{}
	Next *NodeL
}

// List represents the linked list structure
type List struct {
	Head *NodeL
	Tail *NodeL
}

// ListPushFront inserts a new element at the beginning of the list
func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}
	if l.Head == nil {
		// If the list is empty, newNode becomes both Head and Tail
		l.Head = newNode
		l.Tail = newNode
	} else {
		// Otherwise, insert newNode at the front
		newNode.Next = l.Head
		l.Head = newNode
	}
}
