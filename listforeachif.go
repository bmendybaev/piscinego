package piscine

// IsPositiveNode checks if the node contains a positive integer
func IsPositiveNode(node *NodeL) bool {
	switch v := node.Data.(type) {
	case int:
		return v > 0
	case float32:
		return v > 0
	case float64:
		return v > 0
	case byte:
		return v > 0
	default:
		return false
	}
}

// IsAlNode checks if the node contains a non-numeric (string) value
func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}

// ListForEachIf applies function f to nodes that meet the condition cond
func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	current := l.Head
	for current != nil {
		if cond(current) {
			f(current)
		}
		current = current.Next
	}
}
