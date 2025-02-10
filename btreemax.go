package piscine

// MaxTreeNode represents a node in the binary tree for the max function.
type MaxTreeNode struct {
	Data  string
	Left  *MaxTreeNode
	Right *MaxTreeNode
}

// BTreeMax returns the node with the maximum value in the binary search tree.
func BTreeMax(root *MaxTreeNode) *MaxTreeNode {
	if root == nil {
		return nil
	}
	// Traverse to the rightmost node
	current := root
	for current.Right != nil {
		current = current.Right
	}
	return current
}
