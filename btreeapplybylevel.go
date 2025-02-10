package piscine

func CheckIfBinaryTree(root *TreeNode) bool {
	return validateBST(root, "", "")
}

func validateBST(node *TreeNode, min, max string) bool {
	if node == nil {
		return true
	}

	nodeData := node.Data

	if (min != "" && nodeData <= min) || (max != "" && nodeData >= max) {
		return false
	}

	return validateBST(node.Left, min, nodeData) && validateBST(node.Right, nodeData, max)
}
