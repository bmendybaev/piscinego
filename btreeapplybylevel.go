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

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			f(node.Data)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
}
