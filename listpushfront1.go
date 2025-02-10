package piscine

// NodeL represents an element in the linked list
type NodeL struct {
	Data interface{}
	Next *NodeL
}

// List represents the linked list with head and tail pointers
type List struct {
	Head *NodeL
	Tail *NodeL
}

// ListPushFront inserts a new node with the given data at the beginning of the list
func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}
	
	// If the list is empty, both Head and Tail point to the new node
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		// Insert the new node at the front
		newNode.Next = l.Head
		l.Head = newNode
	}
}