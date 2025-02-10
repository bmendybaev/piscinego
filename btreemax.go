package piscine  

// TreeNode represents a node in the binary tree.  
type TreeNode struct {  
	Data  string  
	Left  *TreeNode  
	Right *TreeNode  
}  

// BTreeMax returns the node with the maximum value in the binary search tree.  
func BTreeMax(root *TreeNode) *TreeNode {  
	if root == nil {  
		return nil  
	}  
	// Traverse to the rightmost node  
	current := root  
	for current.Right != nil {  
		current = current.Right  
	}  
	return current  
}
