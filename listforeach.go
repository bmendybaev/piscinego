package piscine

// ListForEach applies the function f to each node in the list
func ListForEach(l *List, f func(*NodeL)) {
	current := l.Head
	for current != nil {
		f(current)
		current = current.Next
	}
}

// Add2_node adds 2 to integers and appends "2" to strings in the node's Data
func Add2_node(node *NodeL) {
	switch v := node.Data.(type) {
	case int:
		node.Data = v + 2
	case string:
		node.Data = v + "2"
	}
}

// Subtract3_node subtracts 3 from integers and appends "-3" to strings in the node's Data
func Subtract3_node(node *NodeL) {
	switch v := node.Data.(type) {
	case int:
		node.Data = v - 3
	case string:
		node.Data = v + "-3"
	}
}
