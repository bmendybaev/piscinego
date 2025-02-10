package piscine

// BTreeSearchItem returns the TreeNode with a data field equal to elem if it exists in the tree, otherwise returns nil
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Data == elem {
		return root
	}

	if left := BTreeSearchItem(root.Left, elem); left != nil {
		return left
	}

	return BTreeSearchItem(root.Right, elem)
}

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
