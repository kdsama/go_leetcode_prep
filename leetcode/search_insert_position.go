package leetcode

import "fmt"

func SearchInsert(nums []int, target int) int {

	start := 0
	end := len(nums) - 1
	median := 0

	for start <= end {
		median = (start + end) / 2

		if nums[median] > target {
			end = median - 1

		} else if target > nums[median] {
			start = median + 1

		} else {

			return median
		}
	}
	fmt.Println(median, "<<<<<<<<<<<<<<<")
	if nums[median] > target {
		return median
	} else {
		return median + 1
	}
}
