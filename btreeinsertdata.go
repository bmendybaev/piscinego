package piscine

// TreeNode defines the structure of a node in the binary search tree
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data               string
}

// BTreeInsertData inserts new data into the binary search tree following BST properties
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Left, data)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Right, data)
		}
	}

	return root
}
