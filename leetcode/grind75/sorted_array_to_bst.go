package grind75

import "algorithms.com/m/leetcode/data_structures"

func SortedArrayToBST(nums []int) *data_structures.TreeNode {

	// height balanced
	// something to do with the central position it seems

	len_nums := len(nums)
	if len_nums == 0 {
		return nil
	}
	mid := len_nums / 2
	root := &data_structures.TreeNode{nums[mid], nil, nil}
	if mid == 0 {
		return root
	}
	root.Left = SortedArrayToBST(nums[:mid])
	root.Right = SortedArrayToBST(nums[mid+1:])
	return root

}
