package piscine

type TreeNode struct {
	Data   string
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

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

func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, "", "")
}

func isBST(node *TreeNode, min, max string) bool {
	if node == nil {
		return true
	}

	nodeData := node.Data

	if (min != "" && nodeData <= min) || (max != "" && nodeData >= max) {
		return false
	}

	return isBST(node.Left, min, nodeData) && isBST(node.Right, nodeData, max)
}

func BTreeApplyByLevel(root *TreeNode, f func(string, int)) {
	if root == nil {
		return
	}

	type NodeLevel struct {
		node  *TreeNode
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
