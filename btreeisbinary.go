package piscine

// BTreeLevelCount returns the number of levels of the binary tree (height of the tree)
func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := BTreeLevelCount(root.Left)
	rightHeight := BTreeLevelCount(root.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// BTreeIsBinary returns true only if the tree given by root follows the binary search tree properties
func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, nil, nil)
}

func isBST(node *TreeNode, min, max *string) bool {
	if node == nil {
		return true
	}

	nodeData := node.Data.(string)
	if (min != nil && nodeData <= *min) || (max != nil && nodeData >= *max) {
		return false
	}

	return isBST(node.Left, min, &nodeData) && isBST(node.Right, &nodeData, max)
}
