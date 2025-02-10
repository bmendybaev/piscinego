package main

import (
	"fmt"
)

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

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Data == elem {
		return root
	}

	if left := BTreeSearchItem(root.Left, elem); left != nil {
		return left
	}

	return BTreeSearchItem(root.Right, elem)
}

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := BTreeLevelCount(root.Left)
	rightHeight := BTreeLevelCount(root.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")

	fmt.Println("Is Binary Search Tree:", BTreeIsBinary(root))

	selected := BTreeSearchItem(root, "7")
	fmt.Print("Item selected -> ")
	if selected != nil {
		fmt.Println(selected.Data)
	} else {
		fmt.Println("nil")
	}

	fmt.Print("Parent of selected item -> ")
	if selected != nil && selected.Parent != nil {
		fmt.Println(selected.Parent.Data)
	} else {
		fmt.Println("nil")
	}

	fmt.Print("Left child of selected item -> ")
	if selected != nil && selected.Left != nil {
		fmt.Println(selected.Left.Data)
	} else {
		fmt.Println("nil")
	}

	fmt.Print("Right child of selected item -> ")
	if selected != nil && selected.Right != nil {
		fmt.Println(selected.Right.Data)
	} else {
		fmt.Println("nil")
	}

	fmt.Println("Tree Level Count:", BTreeLevelCount(root))
}
