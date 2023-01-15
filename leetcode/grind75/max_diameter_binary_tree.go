package grind75

import "algorithms.com/m/leetcode/data_structures"

func diameterOfBinaryTree(root *data_structures.TreeNode) int {
	max_diameter := 0
	if root == nil {
		return 0
	}
	maxDepth(root, &max_diameter)

	return max_diameter
}

func maxDepth(root *data_structures.TreeNode, max_diameter *int) int {
	if root == nil {
		return 0
	}

	var leftRoot = maxDepth(root.Left, max_diameter)
	var rightRoot = maxDepth(root.Right, max_diameter)
	*max_diameter = max(leftRoot+rightRoot, (*max_diameter))
	if leftRoot > rightRoot {
		return leftRoot + 1
	}
	return rightRoot + 1
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
