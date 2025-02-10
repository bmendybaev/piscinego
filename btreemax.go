package piscine

// TreeNode structure for the binary tree
type TreeNodeMMM struct {
	Data  string
	Left  *TreeNodeMMM
	Right *TreeNodeMMM
}

// BTreeInsertData function to insert data into the tree
func BTreeInsertDataMAX(root *TreeNodeMMM, data string) {
	if root == nil {
		return
	}
	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNodeMMM{Data: data}
		} else {
			BTreeInsertDataMAX(root.Left, data)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNodeMMM{Data: data}
		} else {
			BTreeInsertDataMAX(root.Right, data)
		}
	}
}

// BTreeMax function to find the node with the maximum value
func BTreeMax(root *TreeNodeMMM) *TreeNodeMMM {
	if root == nil {
		return nil
	}
	if root.Right != nil {
		return BTreeMax(root.Right) // Continue to right until you can’t
	}
	return root // The maximum value node
}
