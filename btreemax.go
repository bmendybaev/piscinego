package piscine  

// TreeNode structure for the binary tree  
type TreeNode struct {  
    Data  string  
    Left  *TreeNode  
    Right *TreeNode  
}  

// BTreeInsertData function to insert data into the tree  
func BTreeInsertData(root *TreeNode, data string) {  
    if root == nil {  
        return  
    }  
    if data < root.Data {  
        if root.Left == nil {  
            root.Left = &TreeNode{Data: data}  
        } else {  
            BTreeInsertData(root.Left, data)  
        }  
    } else {  
        if root.Right == nil {  
            root.Right = &TreeNode{Data: data}  
        } else {  
            BTreeInsertData(root.Right, data)  
        }  
    }  
}  

// BTreeMax function to find the node with the maximum value  
func BTreeMax(root *TreeNode) *TreeNode {  
    if root == nil {  
        return nil  
    }  
    if root.Right != nil {  
        return BTreeMax(root.Right) // Continue to right until you can’t  
    }  
    return root // The maximum value node  
}
