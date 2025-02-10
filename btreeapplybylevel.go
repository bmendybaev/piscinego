package piscine

func CheckIfBinaryTree(root *TreeNode) bool {
	return validateBST(root, "", "")
}

func validateBST(node *TreeNode, min, max string) bool {
	if node == nil {
		return true
	}

	nodeData := node.Data

	if (min != "" && nodeData <= min) || (max != "" && nodeData >= max) {
		return false
	}

	return validateBST(node.Left, min, nodeData) && validateBST(node.Right, nodeData, max)
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
