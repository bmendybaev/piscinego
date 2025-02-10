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

func BTreeApplyByLevel(root *BinaryTreeNode, f func(string, int)) {
	if root == nil {
		return
	}

	type NodeLevel struct {
		node  *BinaryTreeNode
		level int
	}

	queue := []NodeLevel{{node: root, level: 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		f(current.node.Data, current.level)

		if current.node.Left != nil {
			queue = append(queue, NodeLevel{node: current.node.Left, level: current.level + 1})
		}
		if current.node.Right != nil {
			queue = append(queue, NodeLevel{node: current.node.Right, level: current.level + 1})
		}
	}
}
