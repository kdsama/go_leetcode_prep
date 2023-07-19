package grind75

func search(nums []int, target int) int {

	// log n means we have to do a variation of binary search
	// we probably would need to pass index pointer
	// we dont know where the shift is
	// but we can do some type of derivations
	// bw middle left and right
	// if left > target , we need to go on the right side
	left := 0
	right := len(nums) - 1
	index := -1
	if len(nums) == 1 && nums[0] == target {
		return 0
	}
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		}
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		// first need to figure out the left and the right
		if nums[left] <= nums[middle] {
			// left is greater than middle
			// means a rotation is present
			// left still would be the smallest number till some number
			// if target > middle
			// if
			if target < nums[left] || target > nums[middle] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			if target < nums[middle] || target > nums[right] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return index
}
