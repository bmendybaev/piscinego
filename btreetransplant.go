package piscine

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Data   string
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

// BTreeTransplant replaces the subtree rooted at node with the subtree rooted at rplc
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == root {
		return rplc
	}

	if node.Parent != nil {
		if node.Parent.Left == node {
			node.Parent.Left = rplc
		} else {
			node.Parent.Right = rplc
		}
	}

	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}
