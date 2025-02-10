package piscine

// BTreeTransplant replaces the subtree rooted at node with the subtree rooted at rplc
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	// If node to replace is the root
	if root == node {
		return rplc
	}

	// Find the parent of the node to be replaced
	var parent *TreeNode
	current := root
	for current != nil && current != node {
		parent = current
		if node.Data < current.Data {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	// Replace node with rplc
	if parent != nil {
		if parent.Left == node {
			parent.Left = rplc
		} else {
			parent.Right = rplc
		}
	}

	return root
}
