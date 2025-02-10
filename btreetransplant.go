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

// BTreeTransplant replaces the subtree rooted at node with the subtree rooted at rplc
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	if root == node {
		return rplc
	}

	if node.Data < root.Data {
		root.Left = BTreeTransplant(root.Left, node, rplc)
	} else if node.Data > root.Data {
		root.Right = BTreeTransplant(root.Right, node, rplc)
	}

	return root
}
