package piscine

// BTreeIsBinary returns true only if the tree given by root follows the binary search tree properties
func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, "", "")
}

func isBST(node *TreeNode, min, max string) bool {
	if node == nil {
		return true
	}

	nodeData := node.Data // node.Data is assumed to be of type string

	if (min != "" && nodeData <= min) || (max != "" && nodeData >= max) {
		return false
	}

	return isBST(node.Left, min, nodeData) && isBST(node.Right, nodeData, max)
}
