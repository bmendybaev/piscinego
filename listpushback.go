package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}
	if l.Head == nil {
		// List is empty, set both Head and Tail to the new node
		l.Head = newNode
		l.Tail = newNode
	} else {
		// List is not empty, append to the Tail
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
