package grind75

import (
	"algorithms.com/m/leetcode/data_structures"
)

// given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

func levelOrder(root *data_structures.TreeNode) [][]int {

	// We need a recursion where
	// whenever we go a step down in the three
	// we have the depth of the tree at that level
	// and we check if there is an item in the list = level + 1 of whatever the current level is
	// if not , then append , else append inside

	to_return := [][]int{}
	traverse(root.Left, to_return, 2)
	traverse(root.Right, to_return, 2)

	return to_return
}

func traverse(node *data_structures.TreeNode, to_return [][]int, depth int) {
	if node == nil {
		return

	}
	to_return[depth] = append(to_return[depth], node.Val)
	traverse(node.Left, to_return, depth+1)
	traverse(node.Right, to_return, depth+1)
}
