package piscine

// BTreeMax returns the node with the maximum value in the binary tree
func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	current := root
	for current.Right != nil {
		current = current.Right
	}
	return current
}

// BTreeMin returns the node with the minimum value in the binary tree
func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	current := root
	for current.Left != nil {
		current = current.Left
	}
	return current
}
