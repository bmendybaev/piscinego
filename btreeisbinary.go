package piscine

// BTreeIsBinary returns true only if the tree given by root follows the binary search tree properties
func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, nil, nil)
}

func isBST(node *TreeNode, min, max *string) bool {
	if node == nil {
		return true
	}

	nodeData, ok := node.Data.(string)
	if !ok {
		return false
	}

	if (min != nil && nodeData <= *min) || (max != nil && nodeData >= *max) {
		return false
	}

	return isBST(node.Left, min, &nodeData) && isBST(node.Right, &nodeData, max)
}
