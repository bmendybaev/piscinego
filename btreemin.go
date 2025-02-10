package piscine

// BTreeMin returns the node with the minimum value in the binary tree
func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	minNode := root

	leftMin := BTreeMin(root.Left)
	rightMin := BTreeMin(root.Right)

	if leftMin != nil && leftMin.Data < minNode.Data {
		minNode = leftMin
	}

	if rightMin != nil && rightMin.Data < minNode.Data {
		minNode = rightMin
	}

	return minNode
}
