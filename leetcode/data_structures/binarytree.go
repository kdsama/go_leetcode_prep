package data_structures

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DisplayBST(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	DisplayBST(root.Left)

	DisplayBST(root.Right)

}
