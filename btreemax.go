package piscine

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

// BTreeInsertData inserts data into the binary search tree
func BTreeInsertData(root **TreeNode, data string) {
	if *root == nil {
		*root = &TreeNode{Data: data}
		return
	}
	if data < (*root).Data {
		BTreeInsertData(&(*root).Left, data)
	} else {
		BTreeInsertData(&(*root).Right, data)
	}
}

// BTreeMax returns the node with the maximum value in the binary tree
func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	current := root
	for current.Right != nil {
		current = current.Right
	}
	return current
}
