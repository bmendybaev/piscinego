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
