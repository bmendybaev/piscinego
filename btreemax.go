package piscine

// TreeNode represents a node in a binary tree
type TreeNodeМ struct {
	Data  string
	Left  *TreeNodeМ
	Right *TreeNodeМ
}

// BTreeInsertData inserts data into the binary search tree
func BTreeInsertData(root **TreeNodeМ, data string) {
	if *root == nil {
		*root = &TreeNodeМ{Data: data}
		return
	}
	if data < (*root).Data {
		BTreeInsertData(&(*root).Left, data)
	} else {
		BTreeInsertData(&(*root).Right, data)
	}
}

// BTreeMax returns the node with the maximum value in the binary tree
func BTreeMax(root *TreeNodeМ) *TreeNodeМ {
	if root == nil {
		return nil
	}

	current := root
	for current.Right != nil {
		current = current.Right
	}
	return current
}
