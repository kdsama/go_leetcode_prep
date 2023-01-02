package grind75

import "fmt"

// Given an array of integers nums which is sorted in ascending
// order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

// You must write an algorithm with O(log n) runtime complexity.

func Search(nums []int, target int) int {

	start := 0
	mid := len(nums) / 2
	end := len(nums) - 1
	if len(nums) == 1 {
		if nums[0] != target {
			return -1
		}
		return 0
	}

	for start <= end {
		mid = (start + end) / 2
		fmt.Println(mid, nums[mid], end)
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			start = mid + 1
			// mid = (mid + end) / 2
		} else {
			end = mid - 1
			// are we here
			// mid = (start + mid) / 2
		}
	}
	return -1
}
