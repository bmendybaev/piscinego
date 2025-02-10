package piscine

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

// BTreeInsertData inserts data into the binary tree
func BTreeInsertData(root *TreeNode, data string) {
	if root == nil {
		return
	}
	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{Data: data}
		} else {
			BTreeInsertData(root.Left, data)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data}
		} else {
			BTreeInsertData(root.Right, data)
		}
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
