package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, "", "")
}

func isBST(node *TreeNode, min, max string) bool {
	if node == nil {
		return true
	}

	nodeData := node.Data

	if (min != "" && nodeData <= min) || (max != "" && nodeData >= max) {
		return false
	}

	return isBST(node.Left, min, nodeData) && isBST(node.Right, nodeData, max)
}
