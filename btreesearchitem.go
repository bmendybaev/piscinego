package piscine

// BTreeApplyPreorder applies the function f to each element in the tree using preorder traversal
func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	f(root.Data)
	BTreeApplyPreorder(root.Left, f)
	BTreeApplyPreorder(root.Right, f)
}

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
