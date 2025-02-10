package piscine

type BinaryTreeNode struct {
	Data   string
	Left   *BinaryTreeNode
	Right  *BinaryTreeNode
	Parent *BinaryTreeNode
}

func InsertData(root *BinaryTreeNode, data string) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{Data: data}
	}
	if data < root.Data {
		if root.Left == nil {
			root.Left = &BinaryTreeNode{Data: data, Parent: root}
		} else {
			InsertData(root.Left, data)
		}
	} else {
		if root.Right == nil {
			root.Right = &BinaryTreeNode{Data: data, Parent: root}
		} else {
			InsertData(root.Right, data)
		}
	}
	return root
}

func IsBinarySearchTree(root *BinaryTreeNode) bool {
	return checkBST(root, "", "")
}

func checkBST(node *BinaryTreeNode, min, max string) bool {
	if node == nil {
		return true
	}

	nodeData := node.Data

	if (min != "" && nodeData <= min) || (max != "" && nodeData >= max) {
		return false
	}

	return checkBST(node.Left, min, nodeData) && checkBST(node.Right, nodeData, max)
}

func IsBinarySearchTreeOutput(root *BinaryTreeNode) string {
	if IsBinarySearchTree(root) {
		return "true"
	}
	return "false"
}
