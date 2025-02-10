package piscine

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
