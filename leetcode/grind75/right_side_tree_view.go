package grind75

import (
	"fmt"

	"algorithms.com/m/leetcode/data_structures"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = data_structures.TreeNode

func rightSideView(root *TreeNode) []int {
	d := depth(root)
	v := make([]int, d)
	for i := 0; i < d; i++ {
		v[i] = -1
	}
	traversal(root, 0, &v)
	return v

}
func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := depth(node.Left)
	r := depth(node.Right)
	if l > r {
		return 1 + l
	}
	return 1 + r
}

func traversal(node *TreeNode, depth int, arr *[]int) {
	a := *arr

	// we need to figure out this better
	if depth >= len(a) {
		return
	}

	if node == nil {
		return
	}
	fmt.Println(node.Val)
	if a[depth] == -1 {
		a[depth] = node.Val
	}

	if node.Right != nil {
		traversal(node.Right, depth+1, arr)
	}
	traversal(node.Left, depth+1, arr)

}
