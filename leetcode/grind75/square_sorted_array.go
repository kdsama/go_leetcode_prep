package grind75

func SortedSquares(nums []int) []int {

	// O(n) solution , how
	// find the first positive number
	if len(nums) == 1 {
		nums[0] = nums[0] * nums[0]
		return nums
	}
	i := 0
	for i = 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			break
		}
	}

	new_arr := []int{}

	if i == 0 {

		for i := range nums {
			new_arr = append(new_arr, nums[i]*nums[i])
		}
		return new_arr

	} else if i == len(nums) {
		for i := len(nums) - 1; i >= 0; i-- {
			new_arr = append(new_arr, nums[i]*nums[i])
		}
		return new_arr
	}
	left := i - 1
	right := i

	for left < right && left >= 0 && right < len(nums) {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			new_arr = append(new_arr, nums[left]*nums[left])
			left--
		} else {
			new_arr = append(new_arr, nums[right]*nums[right])
			right++
		}
	}
	for left >= 0 {
		new_arr = append(new_arr, nums[left]*nums[left])
		left--
	}
	for right < len(nums) {
		new_arr = append(new_arr, nums[right]*nums[right])
		right++
	}
	return new_arr
}
